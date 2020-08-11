package orionclient

import (
	"context"
	"errors"
	"net/http"
)

func (c *Client) GetVersion(ctx context.Context) (*Version, error) {
	headers := map[string]string{}
	req, err := c.newRequest(ctx, http.MethodGet, "/version", nil, headers, nil)
	if err != nil {
		return nil, err
	}

	var version *Version
	resp, err := c.doRequest(req, &version)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return version, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	default:
		return nil, errors.New("unexpected error")
	}
}
