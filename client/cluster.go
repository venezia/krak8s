// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": cluster Resource Client
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateClusterPath computes a request path to the create action of cluster.
func CreateClusterPath(project string, ns string) string {
	param0 := project
	param1 := ns

	return fmt.Sprintf("/v1/projects/%s/ns/%s/cluster", param0, param1)
}

// Create the specified cluster resources
func (c *Client) CreateCluster(ctx context.Context, path string, payload *CluterPostBody, contentType string) (*http.Response, error) {
	req, err := c.NewCreateClusterRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateClusterRequest create the request corresponding to the create action endpoint of the cluster resource.
func (c *Client) NewCreateClusterRequest(ctx context.Context, path string, payload *CluterPostBody, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteClusterPath computes a request path to the delete action of cluster.
func DeleteClusterPath(project string, ns string) string {
	param0 := project
	param1 := ns

	return fmt.Sprintf("/v1/projects/%s/ns/%s/cluster", param0, param1)
}

// Delete the cluster resource
func (c *Client) DeleteCluster(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteClusterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteClusterRequest create the request corresponding to the delete action endpoint of the cluster resource.
func (c *Client) NewDeleteClusterRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetClusterPath computes a request path to the get action of cluster.
func GetClusterPath(project string, ns string) string {
	param0 := project
	param1 := ns

	return fmt.Sprintf("/v1/projects/%s/ns/%s/cluster", param0, param1)
}

// Get the status of the cluster resources
func (c *Client) GetCluster(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetClusterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetClusterRequest create the request corresponding to the get action endpoint of the cluster resource.
func (c *Client) NewGetClusterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}