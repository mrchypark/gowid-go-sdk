package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestClient points a client at a server that always answers with status and
// the given literal body.
func newTestClient(t *testing.T, apiKey string, status int, body string) *Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	c := NewClient(apiKey)
	c.BaseURL = server.URL
	return c
}

func TestDoSuccessEnvelope(t *testing.T) {
	// Literal official members payload, independent of the request DTOs.
	body := `{
  "result": { "code": 20000000, "desc": "success" },
  "totalCount": 2,
  "data": [
    {
      "userId": 9007199254740,
      "userName": "홍길동",
      "email": "hong@example.com",
      "isContractor": false,
      "status": "ACTIVE",
      "department": { "id": 12, "name": "개발팀" },
      "position": "매니저",
      "role": { "type": "ADMIN", "name": "관리자", "description": null },
      "notificationOnOff": true
    },
    {
      "userId": 9007199254741,
      "userName": "이순신",
      "email": "lee@example.com",
      "isContractor": true,
      "status": "ACTIVE",
      "department": { "id": 13, "name": "영업팀" },
      "position": "대리",
      "role": { "type": "USER", "name": "사용자", "description": "일반 사용자" },
      "notificationOnOff": false
    }
  ]
}`
	c := newTestClient(t, "secret-key", http.StatusOK, body)

	resp, err := c.GetMembers()
	if err != nil {
		t.Fatalf("GetMembers: %v", err)
	}
	if resp.Result.Code != SuccessCode {
		t.Fatalf("result code = %d, want %d", resp.Result.Code, SuccessCode)
	}
	if resp.TotalCount != 2 {
		t.Errorf("totalCount = %d, want 2", resp.TotalCount)
	}
	if len(resp.Data) != 2 {
		t.Fatalf("members = %d, want 2", len(resp.Data))
	}
	// The official schema is int64; a value past 2^53 must not be truncated.
	if resp.Data[0].UserID != 9007199254740 {
		t.Errorf("userId = %d, want 9007199254740", resp.Data[0].UserID)
	}
	if resp.Data[1].Role.Description != "일반 사용자" {
		t.Errorf("role description = %q", resp.Data[1].Role.Description)
	}
}

func TestDoServiceErrorOnHTTPStatus(t *testing.T) {
	// Literal official error envelope: 401 + result 40100010.
	c := newTestClient(t, "bad-key", http.StatusUnauthorized,
		`{"result":{"code":40100010,"desc":"유효하지 않은 API Key"},"data":null}`)

	_, err := c.GetMembers()
	if err == nil {
		t.Fatal("expected an error for 40100010")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *APIError", err)
	}
	if apiErr.StatusCode != http.StatusUnauthorized {
		t.Errorf("StatusCode = %d, want 401", apiErr.StatusCode)
	}
	if apiErr.Code != 40100010 {
		t.Errorf("Code = %d, want 40100010", apiErr.Code)
	}
	if apiErr.Desc != "유효하지 않은 API Key" {
		t.Errorf("Desc = %q", apiErr.Desc)
	}
	if !strings.Contains(err.Error(), "40100010") {
		t.Errorf("Error() = %q, want it to mention the result code", err.Error())
	}
}

func TestDoNonJSONErrorKeepsRawBody(t *testing.T) {
	c := newTestClient(t, "secret-key", http.StatusBadGateway, "<html>502 Bad Gateway</html>")

	_, err := c.GetMembers()
	if err == nil {
		t.Fatal("expected an error for a non-JSON 502")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *APIError", err)
	}
	if apiErr.Code != 0 {
		t.Errorf("Code = %d, want 0 for a non-envelope body", apiErr.Code)
	}
	if !strings.Contains(apiErr.Body, "502 Bad Gateway") {
		t.Errorf("Body = %q, want the raw payload", apiErr.Body)
	}
	if !strings.Contains(err.Error(), "502 Bad Gateway") {
		t.Errorf("Error() = %q, want it to fall back to the raw body", err.Error())
	}
}

func TestDoRejectsNonSuccessCodeOn200(t *testing.T) {
	// HTTP 200 with a failure result code must still be an error.
	c := newTestClient(t, "secret-key", http.StatusOK,
		`{"result":{"code":40020020,"desc":"해당 ID로 이용내역을 찾을 수 없음"},"data":null}`)

	active := true
	_, err := c.GetPurposesV2(&active)
	if err == nil {
		t.Fatal("expected an error for a failure code on HTTP 200")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *APIError", err)
	}
	if apiErr.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want 200", apiErr.StatusCode)
	}
	if apiErr.Code != 40020020 {
		t.Errorf("Code = %d, want 40020020", apiErr.Code)
	}
}

func TestDoPrefersEnvelopeErrorOverDecodeError(t *testing.T) {
	// Failure envelope whose data shape cannot decode into the target. The
	// result code must win, not a json type error.
	c := newTestClient(t, "secret-key", http.StatusOK,
		`{"result":{"code":40000001,"desc":"잘못된 파라미터"},"data":"unexpected-string"}`)

	active := true
	_, err := c.GetPurposesV2(&active)
	if err == nil {
		t.Fatal("expected an error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T (%v), want *APIError", err, err)
	}
	if apiErr.Code != 40000001 {
		t.Errorf("Code = %d, want 40000001", apiErr.Code)
	}
	if apiErr.Desc != "잘못된 파라미터" {
		t.Errorf("Desc = %q", apiErr.Desc)
	}
}

func TestDoRejectsMalformedAndMissingResult(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{name: "not json", body: `not json at all`},
		{name: "truncated json", body: `{"result":{"code":20000000,`},
		{name: "missing result", body: `{"totalCount":0,"data":[]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newTestClient(t, "secret-key", http.StatusOK, tt.body)
			if _, err := c.GetMembers(); err == nil {
				t.Fatalf("expected an error for body %q", tt.body)
			}
		})
	}
}

func TestDoSendsRawAuthorizationKey(t *testing.T) {
	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"result":{"code":20000000,"desc":"success"},"data":[]}`))
	}))
	defer server.Close()

	c := NewClient("plain-api-key")
	c.BaseURL = server.URL
	if _, err := c.GetMembers(); err != nil {
		t.Fatalf("GetMembers: %v", err)
	}
	if gotAuth != "plain-api-key" {
		t.Errorf("Authorization = %q, want the raw key with no Bearer prefix", gotAuth)
	}
}

func TestDoPropagatesTransportError(t *testing.T) {
	c := NewClient("secret-key")
	c.BaseURL = "http://127.0.0.1:1" // nothing listens here

	if _, err := c.GetMembers(); err == nil {
		t.Fatal("expected a transport error")
	}
}

func TestDoRejectsNilHTTPClient(t *testing.T) {
	c := &Client{BaseURL: "https://openapi.gowid.com", APIKey: "secret-key"}

	_, err := c.GetMembers()
	if err == nil {
		t.Fatal("expected an error when HTTPClient is nil")
	}
	if !strings.Contains(err.Error(), "HTTPClient is nil") {
		t.Errorf("Error() = %q, want it to name the nil HTTPClient", err.Error())
	}
}
