package main

import (
	"time"
)

type Expense struct {
	ID          int
	Category    string
	Amount      float64
	Description string
	Date        time.Time
}

func NewExpense(category string, amount float64, description string) *Expense {
	return &Expense{
		Category:    category,
		Amount:      amount,
		Description: description,
		Date:        time.Now(),
	}
}