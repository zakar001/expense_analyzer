package main

import (
	"fmt"
	"strconv"
	"time"
)

type Expense struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

type ExpenseManager struct {
	expenses []Expense
	nextID   int
}

func NewExpenseManager() *ExpenseManager {
	return &ExpenseManager{
		expenses: make([]Expense, 0),
		nextID:   1,
	}
}

func (em *ExpenseManager) AddExpense(amountStr, category, description string) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Error: Invalid amount '%s'\n", amountStr)
		return
	}

	expense := Expense{
		ID:          em.nextID,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now(),
	}

	em.expenses = append(em.expenses, expense)
	em.nextID++

	fmt.Printf("Expense added: $%.2f for %s (%s)\n", amount, category, description)
}

func (em *ExpenseManager) ListAllExpenses() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}

	fmt.Println("All Expenses:")
	fmt.Println("ID | Amount | Category | Description | Date")
	fmt.Println("--------------------------------------------")
	for _, expense := range em.expenses {
		fmt.Printf("%d | $%.2f | %s | %s | %s\n",
			expense.ID,
			expense.Amount,
			expense.Category,
			expense.Description,
			expense.Date.Format("2006-01-02 15:04"))
	}
}

func (em *ExpenseManager) GenerateSummaryReport() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses to report.")
		return
	}

	total := 0.0
	for _, expense := range em.expenses {
		total += expense.Amount
	}

	fmt.Printf("\n=== EXPENSE SUMMARY REPORT ===\n")
	fmt.Printf("Total Expenses: %d\n", len(em.expenses))
	fmt.Printf("Total Amount: $%.2f\n", total)
	fmt.Printf("Average per expense: $%.2f\n", total/float64(len(em.expenses)))
}

func (em *ExpenseManager) GenerateCategoryReport() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses to report.")
		return
	}

	categoryTotals := make(map[string]float64)
	categoryCounts := make(map[string]int)

	for _, expense := range em.expenses {
		categoryTotals[expense.Category] += expense.Amount
		categoryCounts[expense.Category]++
	}

	fmt.Printf("\n=== CATEGORY BREAKDOWN REPORT ===\n")
	fmt.Println("Category | Total Amount | Count | Average")
	fmt.Println("-----------------------------------------")
	
	totalAll := 0.0
	countAll := 0
	
	for category, total := range categoryTotals {
		count := categoryCounts[category]
		average := total / float64(count)
		fmt.Printf("%s | $%.2f | %d | $%.2f\n", 
			category, total, count, average)
		totalAll += total
		countAll += count
	}
	
	fmt.Printf("-----------------------------------------\n")
	fmt.Printf("TOTAL | $%.2f | %d | $%.2f\n", 
		totalAll, countAll, totalAll/float64(countAll))
}