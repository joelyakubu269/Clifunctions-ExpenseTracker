package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadExpense() []Expense {
	data, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		err = os.WriteFile("tasks.json", []byte("[]"), 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return []Expense{}
		}
		data = []byte("[]")
	} else if err != nil {
		fmt.Println("Error:", err)
	}
	var expenses []Expense
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		fmt.Println("Error loading expenses:", err)
	}
	return expenses
}
func saveExpenses(expenses []Expense) {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = os.WriteFile("tasks.json", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
