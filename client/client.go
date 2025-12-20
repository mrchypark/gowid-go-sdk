package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const BaseURL = "https://openapi.gowid.com"

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

func (c *Client) Do(req *http.Request, v interface{}) error {
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return fmt.Errorf("API request failed with status code %d: %s", res.StatusCode, string(bodyBytes))
	}

	if v != nil {
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
