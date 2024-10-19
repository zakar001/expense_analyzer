package main

import (
	"time"
)

type Expense struct {
	ID          int
	Amount      float64
	Category    string
	Description string
	Date        time.Time
}

func NewExpense(amount float64, category, description string) *Expense {
	return &Expense{
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now(),
	}
}