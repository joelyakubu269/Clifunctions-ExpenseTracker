package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func loadExpense() ([]Expense, error) {
	file, err := os.Open("expenses.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var expenses []Expense

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var exp Expense
		if err := json.Unmarshal(scanner.Bytes(), &exp); err != nil {
			return nil, err
		}
		expenses = append(expenses, exp)
	}

	return expenses, scanner.Err()
}

func saveExpenses(expenses []Expense) error {
	file, err := os.Create("expenses.json")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, exp := range expenses {
		data, _ := json.Marshal(exp)
		file.Write(append(data, '\n'))
	}

	return nil
}
