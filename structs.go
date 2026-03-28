package main

import "time"

type Expense struct {
	ID          int       `json:"int"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}
