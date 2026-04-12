package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func AddExpense(description string, amount float64, date time.Time) (Expense, error) {
	if description == "" || amount <= 0 {
		return Expense{}, fmt.Errorf("invalid description or amount")
	}
	id := time.Now().UnixNano()

	expense := Expense{
		ID:          int(id),
		Description: description,
		Amount:      amount,
		Date:        date,
	}
	file, err := os.OpenFile("expenses.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // give you access to perform operations on the file
	// such as append, create the file if it does not exist and to write only
	if err != nil {
		return Expense{}, err
	}
	defer file.Close()
	data, err := json.MarshalIndent(expense, "", " ") // put expense in a json format
	if err != nil {
		return Expense{}, err
	}
	_, err = file.Write(append(data, '\n')) // append data(i.e expense in json format to expenses.json)
	if err != nil {
		return Expense{}, err
	}
	return expense, nil
}
func deleteExpense(id int) error {
	expenses, err := loadExpense()
	if err != nil {
		return err
	}
	var updated []Expense
	found := false
	for _, r := range expenses {
		if r.ID == id {
			found = true
			continue // i have issues on whether to reset the ids after deleting
		}
		updated = append(updated, r)
	}
	if !found {
		return fmt.Errorf("expense with ID %d not found", id)

	}
	return saveExpenses(updated)

}
func updateExpense(description string, amount float64) (Expense, error) {
	expenses, err := loadExpense()
	if err != nil {
		return Expense{}, err
	}

	for i, r := range expenses {
		if r.Description == description {
			expenses[i].Amount = amount // made sure to use expenses[i] just using r wil only b
			// modifying the copy (unused write)

			err = saveExpenses(expenses)
			if err != nil {
				return Expense{}, err
			}

			return expenses[i], nil
		}
	}

	return Expense{}, fmt.Errorf("expense with description '%s' not found", description)
}

func ListExpenses() ([]Expense, error) {
	expenses, err := loadExpense()
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
func TotalExpenses() (float64, error) {
	expenses, err := loadExpense()
	if err != nil {
		return 0, err
	}
	sum := 0.00
	for _, r := range expenses {
		sum += r.Amount
	}
	return sum, err

}
func ExpensesByMonth(month int) (float64, error) {
	expenses, err := loadExpense()
	if err != nil {
		return 0, err
	}
	sum := 0.00
	for _, r := range expenses {
		if int(r.Date.Month()) == month {
			sum += r.Amount
		}
	}
	return sum, err
}
