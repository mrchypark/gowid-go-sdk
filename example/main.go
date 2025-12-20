package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gowid-api-go/client"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	apiKey := os.Getenv("API_KEY")

	c := client.NewClient(apiKey)

	// Test GetMembers
	fmt.Println("Testing GetMembers...")
	members, err := c.GetMembers(&client.GetMembersOptions{Limit: 1})
	if err != nil {
		log.Fatalf("Error getting members: %v", err)
	}
	fmt.Printf("Parsed %d members\n", len(members.Data))
	if len(members.Data) > 0 {
		fmt.Printf("First member: %+v\n", members.Data[0])
	}
	fmt.Println("--------------------------------------------------")

	// Test GetExpenses
	fmt.Println("Testing GetExpenses...")
	// Use a wide date range
	expenses, err := c.GetExpenses(&client.GetExpensesOptions{
		Limit:     1,
		StartDate: "2024-01-01",
		EndDate:   "2025-12-31",
	})
	if err != nil {
		fmt.Printf("Error getting expenses (known issue): %v\n", err)
	} else {
		fmt.Printf("Parsed %d expenses\n", len(expenses.Data))
		if len(expenses.Data) > 0 {
			fmt.Printf("First expense ID: %s\n", expenses.Data[0].ExpenseId)
		}
	}
	fmt.Println("--------------------------------------------------")

	// Test GetPurposes
	fmt.Println("Testing GetPurposes...")
	purposes, err := c.GetPurposes(&client.GetPurposesOptions{Limit: 1})
	if err != nil {
		log.Fatalf("Error getting purposes: %v", err)
	}
	fmt.Printf("Parsed %d purposes\n", len(purposes.Data))
	if len(purposes.Data) > 0 {
		fmt.Printf("First purpose: %+v\n", purposes.Data[0])
	}
	fmt.Println("--------------------------------------------------")
}
