package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Type: [expense-tracker list] to know how to use the cli")
	if len(os.Args) < 2 {
		fmt.Println("Not enough number of arguements")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 6 {
			fmt.Println("not enough number of arguements")
			return
		}
		//description:= os.Args[3]
		description := strings.Join(os.Args[3:4], " ")
		amount := os.Args[5]
		val := ID(amount)
		AddExpense(description, val)
	case "list":
		if len(os.Args) < 2 {
			fmt.Println("Not enough number of arguements")
			return
		}
		ListExpenses()
	case "summary":
		if len(os.Args) < 2 {
			fmt.Println("Not enough number of arguements")
			return
		}
		sum := TotalExpenses()
		fmt.Printf("Total expenses: $%d\n", sum)

	case "delete":
		if len(os.Args) < 4 {
			fmt.Println("Not enough number of arguements")
			return
		}
		val := os.Args[3]
		id := ID(val)
		deleteExpense(id)
	case "monthlyExpense":
		if len(os.Args) < 4 {
			fmt.Println("Not enough number of arguements")
			return
		}
		month := os.Args[2]
		val := os.Args[3]
		id := ID(val)
		sum := ExpensesByMonth(id)
		fmt.Printf("Total expenses for month %d: $%d\n", month, sum)

	}

}
func ID(amount string) int {
	n, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Input a valid number")
	}
	return n
}
