package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const BaseURL = "https://openapi.gowid.com"

// SuccessCode is the official result code for a successful call.
const SuccessCode = 20000000

// APIError reports a failed Gowid call. Code and Desc come from the common
// response envelope; Body keeps the raw payload, used when the response is not
// JSON (gateway errors, HTML error pages) or carries no envelope.
type APIError struct {
	StatusCode int
	Code       int
	Desc       string
	Body       string
}

func (e *APIError) Error() string {
	detail := e.Desc
	if detail == "" {
		detail = strings.TrimSpace(e.Body)
	}
	if e.Code != 0 {
		return fmt.Sprintf("gowid: http %d result %d: %s", e.StatusCode, e.Code, detail)
	}
	if detail == "" {
		return fmt.Sprintf("gowid: http %d", e.StatusCode)
	}
	return fmt.Sprintf("gowid: http %d: %s", e.StatusCode, detail)
}

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// envelopeResult is the smallest view of the official response envelope.
// Result is a pointer so a missing result is distinguishable from code 0.
type envelopeResult struct {
	Result *Result `json:"result"`
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	if c.HTTPClient == nil {
		return fmt.Errorf("gowid: Client.HTTPClient is nil; use NewClient or set your own *http.Client")
	}

	// The official API expects the raw key value, with no Bearer prefix.
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var env envelopeResult
	envErr := json.Unmarshal(body, &env)
	apiErr := &APIError{StatusCode: res.StatusCode, Body: string(body)}
	if envErr == nil && env.Result != nil {
		apiErr.Code = env.Result.Code
		apiErr.Desc = env.Result.Desc
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return apiErr
	}
	// A 2xx status still fails when the envelope is unreadable or the result
	// code is not the official success code. These checks run before decoding
	// the target so a failure envelope never reports a JSON type error instead
	// of its result code.
	if envErr != nil {
		return fmt.Errorf("gowid: malformed response envelope: %w", envErr)
	}
	if env.Result == nil {
		return fmt.Errorf("gowid: response is missing result code: %s", apiErr.Body)
	}
	if env.Result.Code != SuccessCode {
		return apiErr
	}

	if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) newRequest(method, url string, body interface{}) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	return req, nil
}
