package orionclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"strings"
)

func (c *Client) GetRegistrations(ctx context.Context, fs string, fsp string) ([]*Registration, error) {
	headers := map[string]string{
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}

	req, err := c.newRequest(ctx, http.MethodGet, "/v2/registrations", nil, headers, nil)
	if err != nil {
		return nil, err
	}

	var registrations []*Registration
	resp, err := c.doRequest(req, &registrations)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return registrations, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	default:
		return nil, errors.New("unexpected error")
	}
}

func (c *Client) GetRegistration(ctx context.Context, id string, fs string, fsp string) (*Registration, error){
	relativePath := path.Join("/v2/registrations", id)
	headers := map[string]string{
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}
	req, err := c.newRequest(ctx, http.MethodGet, relativePath, nil, headers, nil)
	if err != nil {
		return nil, err
	}

	var registration *Registration
	resp, err := c.doRequest(req, &registration)
	if err != nil { return nil, err }

	switch resp.StatusCode {
	case http.StatusOK:
		return registration, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	default:
		return nil, errors.New("unexpected error")
	}
}

func (c *Client) CreateRegistration(ctx context.Context, registration Registration, fs string, fsp string) (string, error) {
	jsonBytes, _ := json.Marshal(registration)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Fiware-Service": fs,
		"Fiware-ServicePath": fsp,
	}
	req, err := c.newRequest(ctx, http.MethodPost, "/v2/registrations", nil, headers, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}

	resp, err := c.doRequest(req, nil)
	if err != nil {
		return "", err
	}
	locationUrl, _ := resp.Location()
	registrationId := strings.Replace(locationUrl.Path, "/v2/registrations/", "", 1)

	switch resp.StatusCode {
	case http.StatusCreated:
		return registrationId, nil
	case http.StatusBadRequest:
		return "", errors.New("bad request. some parameters may be invalid")
	default:
		return "", errors.New("unexpected error")
	}
}

func (c *Client) DeleteRegistration(ctx context.Context, id string, fs string, fsp string) (error) {
	relativePath := path.Join("/v2/registrations", id)
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
