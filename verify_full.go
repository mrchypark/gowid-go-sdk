// Command verify_full runs a selected set of read-only smoke checks through the
// SDK and reports whether each succeeded against the live API. It covers the
// list endpoints, which need no fabricated IDs; single-resource reads and all
// writes are left out because their result depends on the caller's own data.
//
// It requires a real API_KEY and network access, so it is a manual smoke test,
// not part of `go test`. It exits non-zero if any operation fails.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mrchypark/gowid-go-sdk/client"
)

type check struct {
	name string
	run  func(c *client.Client) error
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "set API_KEY to the key issued by Gowid")
		os.Exit(2)
	}
	c := client.NewClient(apiKey)
	active := true
	start, end := last30Days()

	checks := []check{
		{"GetMembers", func(c *client.Client) error { _, err := c.GetMembers(); return err }},
		{"GetPurposes", func(c *client.Client) error { _, err := c.GetPurposes(nil); return err }},
		{"GetPurposesV2", func(c *client.Client) error { _, err := c.GetPurposesV2(&active); return err }},
		{"GetCardsV2", func(c *client.Client) error {
			_, err := c.GetCardsV2(&client.PageOptionsV2{Page: 0, Size: 20})
			return err
		}},
		{"GetExpenseStatementsV2", func(c *client.Client) error {
			_, err := c.GetExpenseStatementsV2(&client.ExpenseSearchOptionsV2{StartDate: start, EndDate: end, Size: 20})
			return err
		}},
		{"GetNotSubmittedExpensesV2", func(c *client.Client) error {
			_, err := c.GetNotSubmittedExpensesV2(&client.PageOptionsV2{Page: 0, Size: 20})
			return err
		}},
	}

	fmt.Printf("window: %s..%s (KST)\n\n", start, end)
	failed := 0
	for _, k := range checks {
		if err := k.run(c); err != nil {
			fmt.Printf("FAIL %s: %v\n", k.name, err)
			failed++
			continue
		}
		fmt.Printf("ok   %s\n", k.name)
	}
	if failed > 0 {
		fmt.Printf("\n%d of %d operations failed\n", failed, len(checks))
		os.Exit(1)
	}
	fmt.Printf("\nall %d operations succeeded\n", len(checks))
}

// last30Days returns the last 30 days up to today in KST as yyyyMMdd strings.
func last30Days() (start, end string) {
	kst := time.FixedZone("KST", 9*60*60)
	now := time.Now().In(kst)
	return now.AddDate(0, 0, -29).Format("20060102"), now.Format("20060102")
}
