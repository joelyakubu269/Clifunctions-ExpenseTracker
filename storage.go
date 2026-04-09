package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadExpense() ([]Expense, error) {
	data, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		err = os.WriteFile("tasks.json", []byte("[]"), 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to create tasks.json: %v", err)
		}
		data = []byte("[]")
	} else if err != nil {
	return nil, fmt.Errorf("failed to read tasks.json: %v", err)
	}
	var expenses []Expense
	err = json.Unmarshal(data, &expenses)
	return expenses, nil
}
func saveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
