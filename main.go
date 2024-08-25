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

	command := os.Args[1]
	
	switch command {
	case "add":
		if len(os.Args) != 5 {
			fmt.Println("Usage: expense add <amount> <category> <description>")
			return
		}
		amount := os.Args[2]
		category := os.Args[3]
		description := os.Args[4]
		expenseManager.AddExpense(amount, category, description)
		
	case "report":
		if len(os.Args) > 2 && os.Args[2] == "category" {
			expenseManager.GenerateCategoryReport()
		} else {
			expenseManager.GenerateSummaryReport()
		}
		
	case "list":
		expenseManager.ListAllExpenses()
		
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Personal Expense Analyzer")
	fmt.Println("Usage:")
	fmt.Println("  expense add <amount> <category> <description>")
	fmt.Println("  expense list")
	fmt.Println("  expense report")
	fmt.Println("  expense report category")
}