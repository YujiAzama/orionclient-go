package orionclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
	//"strconv"
	"strings"
)

func (c *Client) GetSubscriptions(ctx context.Context, fs string, fsp string) ([]*Subscription, error) {
	//queries := map[string]string{
	//	"limit": strconv.Itoa(limit),
	//}

	headers := map[string]string{
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/v2/subscriptions", nil, headers, nil)
	if err != nil {
		return nil, err
	}

	var subscriptions []*Subscription
	resp, err := c.doRequest(req, &subscriptions)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return subscriptions, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	default:
		return nil, errors.New("unexpected error")
	}
}

func (c *Client) GetSubscription(ctx context.Context, id string, fs string, fsp string) (*Subscription, error){
	relativePath := path.Join("/v2/subscriptions", id)
	headers := map[string]string{
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}
	req, err := c.newRequest(ctx, http.MethodGet, relativePath, nil, headers, nil)
	if err != nil {
		return nil, err
	}

	var subscription *Subscription
	resp, err := c.doRequest(req, &subscription)
	if err != nil { return nil, err }

	switch resp.StatusCode {
	case http.StatusOK:
		return subscription, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	default:
		return nil, errors.New("unexpected error")
	}
}

func (c *Client) CreateSubscription(ctx context.Context, subscription Subscription, fs string, fsp string) (string, error) {
	jsonBytes, _ := json.Marshal(subscription)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}
	req, err := c.newRequest(ctx, http.MethodPost, "/v2/subscriptions", nil, headers, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}

	resp, err := c.doRequest(req, nil)
	if err != nil {
		return "", err
	}
	locationUrl, _ := resp.Location()
	subscriptionId := strings.Replace(locationUrl.Path, "/v2/subscriptions/", "", 1)

	switch resp.StatusCode {
	case http.StatusCreated:
		return subscriptionId, nil
	case http.StatusBadRequest:
		return "", errors.New("bad request. some parameters may be invalid")
	default:
		return "", errors.New("unexpected error")
	}
}

func (c *Client) DeleteSubscription(ctx context.Context, id string, fs string, fsp string) (error) {
	relativePath := path.Join("/v2/subscriptions", id)
	headers := map[string]string{
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}
	req, err := c.newRequest(ctx, http.MethodDelete, relativePath, nil, headers, nil)
	if err != nil {
		return err
	}

	resp, err := c.doRequest(req, nil)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusBadRequest:
		return errors.New("bad request. some parameters may be invalid")
	default:
		return errors.New("unexpected error")
	}
}
