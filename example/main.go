package main

import (
	"fmt"
	"log"
	"os"

	"gowid-api-go/client"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	apiKey := os.Getenv("API_KEY")

	c := client.NewClient(apiKey)

	// Test GetMembers
	fmt.Println("Testing GetMembers...")
	members, err := c.GetMembers(&client.GetMembersOptions{})
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
		Size:      1,
		StartDate: "2024-01-01",
	})
	if err != nil {
		fmt.Printf("Error getting expenses (known issue): %v\n", err)
	} else {
		fmt.Printf("Parsed %d expenses\n", len(expenses.Data.Content))
		if len(expenses.Data.Content) > 0 {
			fmt.Printf("First expense ID: %d\n", expenses.Data.Content[0].ExpenseId)
		}
	}
	fmt.Println("--------------------------------------------------")

	// Test GetExpenses with size/page
	fmt.Println("Testing GetExpenses with size/page...")
	expensesPage, err := c.GetExpenses(&client.GetExpensesOptions{
		Size:      2,
		Page:      1,
		StartDate: "2024-01-01",
	})
	if err != nil {
		fmt.Printf("Error getting expenses page: %v\n", err)
	} else {
		fmt.Printf("Parsed %d expenses on page\n", len(expensesPage.Data.Content))
		if len(expensesPage.Data.Content) > 0 {
			fmt.Printf("Page first expense ID: %d\n", expensesPage.Data.Content[0].ExpenseId)
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
