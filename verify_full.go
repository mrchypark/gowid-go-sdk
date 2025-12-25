package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	baseURL := "https://openapi.gowid.com"

	runTest(apiKey, baseURL, "GET", "/v1/members?limit=1", nil)
	runTest(apiKey, baseURL, "GET", "/v1/purposes?limit=1", nil)
	// Try expenses with dates
	runTest(apiKey, baseURL, "GET", "/v1/expenses?startDate=20251201", nil)
	// Try expenses with dates and page
	runTest(apiKey, baseURL, "GET", "/v1/expenses?startDate=20251201&page=0", nil)

	runTest(apiKey, baseURL, "GET", "/v1/expenses/not-submitted?page=0", nil)

	// Probe for default card
	runTest(apiKey, baseURL, "GET", "/v1/cards", nil)
}

func runTest(apiKey, baseURL, method, path string, body io.Reader) {
	client := &http.Client{Timeout: 10 * time.Second}
	fullURL := baseURL + path
	req, _ := http.NewRequest(method, fullURL, body)
	req.Header.Set("Authorization", apiKey)

	fmt.Printf(">>> Request: %s %s\n", method, fullURL)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("<<< Status: %s\n", resp.Status)
	fmt.Printf("<<< Body: %s\n", string(bodyBytes))

	// Try to parse error code
	var result struct {
		Result struct {
			Code int    `json:"code"`
			Desc string `json:"desc"`
		} `json:"result"`
	}
	json.Unmarshal(bodyBytes, &result)
	if result.Result.Code != 0 && result.Result.Code != 20000000 {
		fmt.Printf("!!! API Error Code: %d, Desc: %s\n", result.Result.Code, result.Result.Desc)
	}
	fmt.Println("--------------------------------------------------")
}
