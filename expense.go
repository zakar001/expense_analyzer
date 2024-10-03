package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Expense struct {
	ID          string    `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type ExpenseManager struct {
	expenses []Expense
	filename string
}

func NewExpenseManager() *ExpenseManager {
	em := &ExpenseManager{
		filename: "expenses.json",
	}
	em.loadExpenses()
	return em
}

func (em *ExpenseManager) AddExpense(amountStr, category, description string) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Error: Invalid amount '%s'\n", amountStr)
		return
	}

	expense := Expense{
		ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now(),
	}

	em.expenses = append(em.expenses, expense)
	em.saveExpenses()
	
	fmt.Printf("Expense added: $%.2f for %s (%s)\n", amount, category, description)
}

func (em *ExpenseManager) GenerateReport() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n=== EXPENSE REPORT ===")
	fmt.Printf("%-12s %-15s %-10s %-20s\n", "Date", "Category", "Amount", "Description")
	fmt.Println("------------------------------------------------------------")
	
	total := 0.0
	categoryTotals := make(map[string]float64)
	
	for _, expense := range em.expenses {
		dateStr := expense.Date.Format("2006-01-02")
		fmt.Printf("%-12s %-15s $%-9.2f %-20s\n", 
			dateStr, expense.Category, expense.Amount, expense.Description)
		
		total += expense.Amount
		categoryTotals[expense.Category] += expense.Amount
	}
	
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Total: $%.2f\n", total)
	
	fmt.Println("\n=== CATEGORY BREAKDOWN ===")
	for category, amount := range categoryTotals {
		percentage := (amount / total) * 100
		fmt.Printf("%-15s: $%.2f (%.1f%%)\n", category, amount, percentage)
	}
}

func (em *ExpenseManager) ShowSummary() {
	if len(em.expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	total := 0.0
	categoryTotals := make(map[string]float64)
	
	for _, expense := range em.expenses {
		total += expense.Amount
		categoryTotals[expense.Category] += expense.Amount
	}
	
	fmt.Println("\n=== EXPENSE SUMMARY ===")
	fmt.Printf("Total Expenses: $%.2f\n", total)
	fmt.Printf("Number of Transactions: %d\n", len(em.expenses))
	fmt.Printf("Average per Transaction: $%.2f\n", total/float64(len(em.expenses)))
	
	fmt.Println("\nSpending by Category:")
	for category, amount := range categoryTotals {
		fmt.Printf("  %-15s: $%.2f\n", category, amount)
	}
}

func (em *ExpenseManager) saveExpenses() {
	file, err := os.Create(em.filename)
	if err != nil {
		fmt.Printf("Error saving expenses: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(em.expenses)
	if err != nil {
		fmt.Printf("Error encoding expenses: %v\n", err)
	}
}

func (em *ExpenseManager) loadExpenses() {
	file, err := os.Open(em.filename)
	if err != nil {
		if os.IsNotExist(err) {
			em.expenses = []Expense{}
			return
		}
		fmt.Printf("Error loading expenses: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&em.expenses)
	if err != nil {
		fmt.Printf("Error decoding expenses: %v\n", err)
		em.expenses = []Expense{}
	}
}