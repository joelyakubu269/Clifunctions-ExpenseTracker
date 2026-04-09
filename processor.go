package main

import (
	"fmt"
	"time"
)

func AddExpense(description string, amount float64) (Expense, error) {
	if description == "" || amount <= 0 {
		return Expense{}, fmt.Errorf("invalid description or amount")
	}
	expenses, err := loadExpense()
	if err != nil {
		return Expense{}, err
	}
	var maxID int

	for _, r := range expenses {
		if r.ID > maxID {
			maxID = r.ID
		}
	}
	maxID++
	expense := Expense{
		ID:          maxID,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	expenses = append(expenses, expense)
	err = saveExpenses(expenses)
	if err != nil {
		return Expense{}, fmt.Errorf("unable to save")
	}
	return expense, nil
}
func deleteExpense(id int) error {
	expenses, err := loadExpense()
	if err != nil {
		return err
	}
	var delete []Expense
	found := false
	for _, r := range expenses {
		if r.ID == id {
			found = true
			continue // i have issues on whether to reset the ids after deleting
		}
		delete = append(delete, r)
	}
	if !found {
		if !found {
			return fmt.Errorf("expense with ID %d not found", id)
		}
	}
	return saveExpenses(delete)

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
			expenses[i].Amount = amount

			err = saveExpenses(expenses)
			if err != nil {
				return Expense{}, err
			}

			return expenses[i], nil
		}
	}

	return Expense{}, fmt.Errorf("expense with description '%s' not found", description)
}

		
		
		


func ListExpenses() ([]Expense,error) {
	expenses, err := loadExpense()
	if err != nil {
		return nil,err
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
