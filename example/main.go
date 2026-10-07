package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mrchypark/gowid-go-sdk/client"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("set API_KEY to the key issued by Gowid")
	}
	c := client.NewClient(apiKey)
	start, end := last30Days()

	members, err := c.GetMembers()
	if err != nil {
		log.Fatalf("GetMembers: %v", err)
	}
	fmt.Printf("members: %d\n", len(members.Data))

	// Dates are yyyyMMdd and both bounds are inclusive.
	statements, err := c.GetExpenseStatementsV2(&client.ExpenseSearchOptionsV2{
		StartDate: start,
		EndDate:   end,
		Page:      0,
		Size:      20,
	})
	if err != nil {
		log.Fatalf("GetExpenseStatementsV2: %v", err)
	}
	fmt.Printf("statements %s..%s: %d\n", start, end, statements.TotalCount)
	for i, s := range statements.Data.Content {
		// Print only non-personal summary fields.
		alias := ""
		if s.CardAlias != nil {
			alias = *s.CardAlias
		}
		fmt.Printf("  %s %d %s %s\n", s.ExpenseDate, s.KRWAmount, s.Currency, alias)
		if i == 4 {
			fmt.Printf("  ... %d more\n", len(statements.Data.Content)-i-1)
			break
		}
	}
}

// last30Days returns the last 30 days up to today in KST as yyyyMMdd strings.
func last30Days() (start, end string) {
	kst := time.FixedZone("KST", 9*60*60)
	now := time.Now().In(kst)
	return now.AddDate(0, 0, -29).Format("20060102"), now.Format("20060102")
}
