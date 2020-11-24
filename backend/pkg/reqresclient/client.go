package reqresclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	get = "GET"
)

// ReqResConfig is the configuration
type ReqResConfig struct {
	baseURL string
	userURL string
	client  iClient
}

type iClient interface {
	Do(*http.Request) (*http.Response, error)
}

// NewReqResConfig constructor
func NewReqResConfig(baseURL string) ReqResConfig {
	return NewReqResConfigWithClient(baseURL, http.DefaultClient)
}

// NewReqResConfigWithClient constructor
func NewReqResConfigWithClient(baseURL string, client iClient) ReqResConfig {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	return ReqResConfig{
		baseURL: baseURL,
		userURL: baseURL + "users/%d",
		client:  client,
	}
}

// HTTPError error
type HTTPError struct {
	resp *http.Response
	body []byte
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("http error: %d %s", e.resp.StatusCode, e.body)
}

// UserResponse response
type UserResponse struct {
	Data User `json:"data"`
}

// User response
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

// GetUser user API
func (c *ReqResConfig) GetUser(ctx context.Context, userID int) (user User, err error) {
	req, err := http.NewRequestWithContext(ctx, get, fmt.Sprintf(c.userURL, userID), nil)
	if err != nil {
		return
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return user, fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, fmt.Errorf("read from io failed: %s", err)
	}

	if code := resp.StatusCode; code < 200 || code > 299 {
		return user, HTTPError{resp: resp, body: body}
	}

	var data UserResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return user, fmt.Errorf("unmarshall failed: %s", err)
	}

	return data.Data, nil
}

func (c *ReqResConfig) doRequest(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.client.Do(req)
	if err != nil {
		return
	}
	return
}
