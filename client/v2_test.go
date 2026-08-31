package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestV2Endpoints(t *testing.T) {
	requests := make([]*http.Request, 0, 13)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.Clone(r.Context()))
		if r.Body != nil {
			body, _ := io.ReadAll(r.Body)
			requests[len(requests)-1].Body = io.NopCloser(strings.NewReader(string(body)))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewClient("secret")
	client.BaseURL = server.URL
	active := true
	amount := int64(1000)
	calls := []func() error{
		func() error { _, err := client.GetExpenseV2(1); return err },
		func() error {
			_, err := client.GetExpenseStatementsV2(&ExpenseSearchOptionsV2{StartDate: "20260101", EndDate: "20260131", Size: 100})
			return err
		},
		func() error { _, err := client.SearchExpensesV2(nil); return err },
		func() error { _, err := client.GetNotSubmittedExpensesV2(&PageOptionsV2{Page: 1}); return err },
		func() error { _, err := client.GetPurposesV2(&active); return err },
		func() error { _, err := client.GetPurposeRequirementsV2(2, 3); return err },
		func() error { _, err := client.UpdateExpenseV2(1, UpdateExpenseRequestV2{ExpenseID: 1}); return err },
		func() error {
			_, err := client.UpdateExpensePurposeV2(1, UpdatePurposeRequestV2{PurposeID: 2})
			return err
		},
		func() error {
			_, err := client.UpdateExpensesPurposeV2(UpdatePurposesRequestV2{ExpenseIDs: []int64{1}, PurposeID: 2})
			return err
		},
		func() error {
			_, err := client.UpdateExpenseParticipantsV2(1, UpdateParticipantsRequestV2{})
			return err
		},
		func() error { _, err := client.UpdateExpenseMemoV2(1, "memo"); return err },
		func() error {
			_, err := client.UpdateExpenseApprovalStatusV2(1, ApproveExpenseRequestV2{ApprovalStatus: "APPROVED", ApprovedAmount: &amount})
			return err
		},
		func() error {
			_, err := client.ApproveExpensesV2([]ApproveExpenseRequestV2{{ExpenseID: 1, ApprovalStatus: "APPROVED"}})
			return err
		},
	}
	for _, call := range calls {
		if err := call(); err != nil {
			t.Fatal(err)
		}
	}
	want := []string{
		"GET /v2/expenses/1", "GET /v2/expense-statements", "GET /v2/expenses", "GET /v2/expenses/not-submitted",
		"GET /v2/purposes", "GET /v2/purposes/2/requirements/3", "PUT /v2/expenses/1", "PUT /v2/expenses/1/purposes",
		"PUT /v2/expenses/purposes", "PUT /v2/expenses/1/participants", "PUT /v2/expenses/1/memo",
		"PUT /v2/expenses/1/approval-status", "PATCH /v2/expenses/approval-status/approved",
	}
	if len(requests) != len(want) {
		t.Fatalf("got %d requests, want %d", len(requests), len(want))
	}
	for i, request := range requests {
		if got := request.Method + " " + request.URL.Path; got != want[i] {
			t.Errorf("request %d = %q, want %q", i, got, want[i])
		}
		if request.Header.Get("Authorization") != "secret" {
			t.Errorf("request %d missing API key", i)
		}
	}
	if got := requests[1].URL.Query().Get("endDate"); got != "20260131" {
		t.Fatalf("endDate = %q", got)
	}
	if got := requests[1].URL.Query().Get("size"); got != "100" {
		t.Fatalf("size = %q", got)
	}
	var purpose UpdatePurposeRequestV2
	if err := json.NewDecoder(requests[7].Body).Decode(&purpose); err != nil || purpose.PurposeID != 2 {
		t.Fatalf("purpose request = %+v, err=%v", purpose, err)
	}
}
