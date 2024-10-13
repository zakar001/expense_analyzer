package main

import (
	"fmt"
	"os"
)

func main() {
	expenseManager := NewExpenseManager()
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: expense-analyzer <command>")
		fmt.Println("Commands: add, report, summary")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: expense-analyzer add <amount> <category> <description>")
			return
		}
		amount := os.Args[2]
		category := os.Args[3]
		description := os.Args[4]
		expenseManager.AddExpense(amount, category, description)
		
	case "report":
		if len(os.Args) > 2 && os.Args[2] == "detailed" {
			expenseManager.GenerateDetailedReport()
		} else {
			expenseManager.GenerateSummaryReport()
		}
		
	case "summary":
		expenseManager.ShowSummary()
		
	default:
		fmt.Println("Unknown command. Available commands: add, report, summary")
	}
}