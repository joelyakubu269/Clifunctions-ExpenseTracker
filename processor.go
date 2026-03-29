package main

import (
	"fmt"
	"time"
)

func AddExpense(description string, amount int) {
	expenses := loadExpense()
	var maxID int
	if len(expenses) == 0 {
		maxID = 1
	}

	for _, r := range expenses {
		if r.ID > maxID {
			maxID = r.ID
		}
	}
	expense := Expense{
		ID:          maxID,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	expenses = append(expenses, expense)
	saveExpenses(expenses)
	fmt.Printf("Expense added successfully (ID: %d)", expense.ID)

}
func deleteExpense(id int) {
	expenses := loadExpense()
	var delete []Expense
	for _, r := range expenses {
		if r.ID != id {
			delete = append(delete, r) // i have issues on whether to reset the ids after deleting
		}
	}
	saveExpenses(delete)
	fmt.Println("Expense deleted successfully")

}
func updateExpense(description string, amount int) {
	expenses := loadExpense()
	found := false
	for i, r := range expenses {
		if r.Description == description {
			expenses[i].Amount = amount // made sure to use expenses[i] just using r wil only b
			// modifying the copy (unused write)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Description:Expense not found")
	}
	saveExpenses(expenses)
	fmt.Printf("The new amount is now %d", amount)
}
func ListExpenses() {
	expenses := loadExpense()
	for _, r := range expenses {
		fmt.Printf("#%d\t%s\t%s\t%d\n",
			r.ID,
			r.Date.Format("2006-01-02"),
			r.Description, r.Amount,
		)
	}
}
func TotalExpenses() int {
	expenses := loadExpense()
	sum := 0
	for _, r := range expenses {
		sum += r.Amount
	}
	return sum

}
func ExpensesByMonth(month int) int {
	expenses := loadExpense()
	sum := 0
	for _, r := range expenses {
		if int(r.Date.Month()) == month {
			sum += r.Amount
		}
	}
	return sum
}
