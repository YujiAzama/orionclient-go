package orionclient

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

type ClientConfig struct {
	Host string
	Port int
	Token string
	TLS bool
	Logger *log.Logger
}

type Client struct {
	BaseURL *url.URL
	HTTPClient *http.Client
	Token string
	Logger *log.Logger
}

func NewClient(config ClientConfig) (*Client, error) {
	schema := "http://"
	if config.TLS {
		schema = "https://"
	}
	baseURL, err := url.Parse(schema + config.Host + ":" + strconv.Itoa(config.Port))
	if err != nil {
		return nil, err
	}
	if config.Logger == nil {
		config.Logger = log.New(os.Stdout, "[LOG]", log.LstdFlags)
	}
	return &Client{
		BaseURL: baseURL,
		HTTPClient: http.DefaultClient,
		Token: config.Token,
		Logger: config.Logger,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, relativePath string, queries, headers map[string]string, reqBody io.Reader) (*http.Request, error) {
	reqURL := *c.BaseURL
	reqURL.Path = path.Join(reqURL.Path, relativePath)

	if queries != nil {
		q := reqURL.Query()
		for k, v := range queries {
			q.Add(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	if reqBody != nil {
		var bodyBuf bytes.Buffer
		_ = io.TeeReader(reqBody, &bodyBuf)
		headers["Content-Length"] = strconv.Itoa(len(bodyBuf.String()))
	} else {
		headers["Content-Length"] = "0"
	}

	if c.Token != "" {
		headers["Authorization"] = "Bearer " + c.Token
	}

	req, err := http.NewRequest(method, reqURL.String(), reqBody)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) doRequest(req *http.Request, respBody interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return resp, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}
	if err := json.Unmarshal(bodyBytes, respBody); err != nil && respBody != nil {
		return resp, err
	}

	return resp, nil
}
