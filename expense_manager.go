package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ExpenseManager struct {
	expenses []*Expense
	nextID   int
}

func NewExpenseManager() *ExpenseManager {
	return &ExpenseManager{
		expenses: make([]*Expense, 0),
		nextID:   1,
	}
}

func (em *ExpenseManager) AddExpense(amountStr, category, description string) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Error: Invalid amount '%s'\n", amountStr)
		return
	}

	expense := NewExpense(amount, category, description)
	expense.ID = em.nextID
	em.nextID++
	em.expenses = append(em.expenses, expense)
	
	fmt.Printf("✓ Added expense: $%.2f for %s (%s)\n", amount, category, description)
}

func (em *ExpenseManager) GenerateSummaryReport() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	categoryTotals := make(map[string]float64)
	totalSpent := 0.0

	for _, expense := range em.expenses {
		categoryTotals[expense.Category] += expense.Amount
		totalSpent += expense.Amount
	}

	fmt.Println("\n=== EXPENSE SUMMARY REPORT ===")
	fmt.Printf("Total expenses: $%.2f\n", totalSpent)
	fmt.Printf("Number of transactions: %d\n", len(em.expenses))
	fmt.Println("\nSpending by category:")
	
	for category, total := range categoryTotals {
		percentage := (total / totalSpent) * 100
		fmt.Printf("  %s: $%.2f (%.1f%%)\n", category, total, percentage)
	}
}

func (em *ExpenseManager) GenerateDetailedReport() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n=== DETAILED EXPENSE REPORT ===")
	total := 0.0
	
	for _, expense := range em.expenses {
		fmt.Printf("#%d: $%.2f | %s | %s | %s\n", 
			expense.ID,
			expense.Amount,
			expense.Category,
			expense.Description,
			expense.Date.Format("2006-01-02 15:04"))
		total += expense.Amount
	}
	
	fmt.Printf("\nTotal: $%.2f\n", total)
}

func (em *ExpenseManager) ShowSummary() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	total := 0.0
	categories := make(map[string]int)
	
	for _, expense := range em.expenses {
		total += expense.Amount
		categories[expense.Category]++
	}

	fmt.Println("\n=== QUICK SUMMARY ===")
	fmt.Printf("Total spent: $%.2f\n", total)
	fmt.Printf("Transactions: %d\n", len(em.expenses))
	fmt.Printf("Categories: %d\n", len(categories))
	
	// Show most expensive category
	maxCategory := ""
	maxAmount := 0.0
	categoryTotals := make(map[string]float64)
	
	for _, expense := range em.expenses {
		categoryTotals[expense.Category] += expense.Amount
		if categoryTotals[expense.Category] > maxAmount {
			maxAmount = categoryTotals[expense.Category]
			maxCategory = expense.Category
		}
	}
	
	if maxCategory != "" {
		fmt.Printf("Highest spending category: %s ($%.2f)\n", maxCategory, maxAmount)
	}
}