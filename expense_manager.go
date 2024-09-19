package main

import (
	"fmt"
	"sort"
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

func (em *ExpenseManager) AddExpense(category string, amount float64, description string) {
	if amount <= 0 {
		fmt.Println("Error: Amount must be positive")
		return
	}

	expense := NewExpense(category, amount, description)
	expense.ID = em.nextID
	em.nextID++
	em.expenses = append(em.expenses, expense)
	
	fmt.Printf("Added expense: $%.2f for %s (%s)\n", amount, description, category)
}

func (em *ExpenseManager) GenerateReport(categoryFilter string) {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n=== EXPENSE REPORT ===")
	if categoryFilter != "" {
		fmt.Printf("Category: %s\n", categoryFilter)
	}
	fmt.Println("ID | Date       | Category | Amount  | Description")
	fmt.Println("---|------------|----------|---------|------------")

	total := 0.0
	count := 0

	for _, expense := range em.expenses {
		if categoryFilter == "" || expense.Category == categoryFilter {
			dateStr := expense.Date.Format("2006-01-02")
			fmt.Printf("%2d | %s | %-8s | $%6.2f | %s\n", 
				expense.ID, dateStr, expense.Category, expense.Amount, expense.Description)
			total += expense.Amount
			count++
		}
	}

	fmt.Printf("\nTotal: $%.2f (%d expenses)\n", total, count)
}

func (em *ExpenseManager) ShowSummary() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n=== SPENDING SUMMARY ===")
	
	// Calculate category totals
	categoryTotals := make(map[string]float64)
	overallTotal := 0.0

	for _, expense := range em.expenses {
		categoryTotals[expense.Category] += expense.Amount
		overallTotal += expense.Amount
	}

	// Sort categories by amount (descending)
	type CategoryTotal struct {
		Category string
		Amount   float64
		Percent  float64
	}
	
	var sortedCategories []CategoryTotal
	for category, amount := range categoryTotals {
		percent := (amount / overallTotal) * 100
		sortedCategories = append(sortedCategories, CategoryTotal{
			Category: category,
			Amount:   amount,
			Percent:  percent,
		})
	}

	sort.Slice(sortedCategories, func(i, j int) bool {
		return sortedCategories[i].Amount > sortedCategories[j].Amount
	})

	fmt.Println("Category    | Amount    | % of Total")
	fmt.Println("------------|-----------|------------")
	for _, ct := range sortedCategories {
		fmt.Printf("%-11s | $%7.2f | %6.1f%%\n", 
			ct.Category, ct.Amount, ct.Percent)
	}
	
	fmt.Printf("\nOverall Total: $%.2f (%d expenses)\n", overallTotal, len(em.expenses))
	
	// Show largest expense
	if len(em.expenses) > 0 {
		largest := em.expenses[0]
		for _, expense := range em.expenses {
			if expense.Amount > largest.Amount {
				largest = expense
			}
		}
		fmt.Printf("Largest expense: $%.2f for %s (%s)\n", 
			largest.Amount, largest.Description, largest.Category)
	}
}