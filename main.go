package main

import (
	"fmt"
	"os"
)

func main() {
	expenseManager := NewExpenseManager()
	
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) != 5 {
			fmt.Println("Usage: go run main.go add <category> <amount> <description>")
			return
		}
		expenseManager.AddExpense(os.Args[2], parseAmount(os.Args[3]), os.Args[4])
		
	case "report":
		if len(os.Args) == 2 {
			expenseManager.GenerateReport("")
		} else {
			expenseManager.GenerateReport(os.Args[2])
		}
		
	case "summary":
		expenseManager.ShowSummary()
		
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Personal Expense Analyzer")
	fmt.Println("Usage:")
	fmt.Println("  go run main.go add <category> <amount> <description>")
	fmt.Println("  go run main.go report [category]")
	fmt.Println("  go run main.go summary")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go add food 25.50 \"Lunch at restaurant\"")
	fmt.Println("  go run main.go report food")
	fmt.Println("  go run main.go summary")
}

func parseAmount(amountStr string) float64 {
	var amount float64
	_, err := fmt.Sscanf(amountStr, "%f", &amount)
	if err != nil {
		fmt.Printf("Error parsing amount: %v\n", err)
		return 0
	}
	return amount
}