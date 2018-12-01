// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Data Client": DataClient Resource Client
//
// Command:
// $ goagen
// --design=DataClient/design
// --out=$(GOPATH)\src\DataClient
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AddDataClientPath computes a request path to the add action of DataClient.
func AddDataClientPath(hash string) string {
	param0 := hash

	return fmt.Sprintf("/data/%s", param0)
}

// add data hash
func (c *Client) AddDataClient(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAddDataClientRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddDataClientRequest create the request corresponding to the add action endpoint of the DataClient resource.
func (c *Client) NewAddDataClientRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DelDataClientPath computes a request path to the del action of DataClient.
func DelDataClientPath(hash string) string {
	param0 := hash

	return fmt.Sprintf("/data/%s", param0)
}

// delete data hash
func (c *Client) DelDataClient(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDelDataClientRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDelDataClientRequest create the request corresponding to the del action endpoint of the DataClient resource.
func (c *Client) NewDelDataClientRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}