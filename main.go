package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Type: [expense-tracker list] to know how to use the cli")
	if len(os.Args) < 2 {
		fmt.Println("Not enough number of arguements")
		fmt.Println("Use 'expense-tracker help' to see available commands")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		//description:= os.Args[3]
		description := addCmd.String("description", "", "add your new expense")
		amount := addCmd.Float64("amount", 0, "the expense amount")
		val := amount
		addCmd.Parse(os.Args[2:])
		addCmd.Usage = func() {
			fmt.Println("Add a new expense")
			fmt.Println("Usage")
			fmt.Println("expense-tracker add -description=\"desc\" -amount=100")
			fmt.Println("\nFlags:")
			addCmd.PrintDefaults()
		}
		AddExpense(*description, *val)
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listCmd.Parse(os.Args[2:])
		listCmd.Usage = func() {
			fmt.Println("Usuage:")
			fmt.Println("expense-tracker list")
			listCmd.PrintDefaults()
		}
		expenses := loadExpense()
		if len(expenses) == 0 {
			fmt.Println("There are no expenses yet")
		}
		ListExpenses()
	case "summary":
		sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		sumCmd.Parse(os.Args[2:])
		sumCmd.Usage = func() {
			fmt.Println("Calculates all your expenses")
			fmt.Println("Usage:")
			fmt.Println("expense-tracker sumarry")
			fmt.Println("\nFlags:")
			sumCmd.PrintDefaults()
		}
		sum := TotalExpenses()
		fmt.Printf("Total expenses: $%.2f\n", sum)

	case "delete":
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := delCmd.Int("id", 0, "the id of the expense you want to delete")
		delCmd.Parse(os.Args[2:])
		delCmd.Usage = func() {
			fmt.Println("Delete an expense")
			fmt.Println("usage:")
			fmt.Println("  expense-tracker delete -id=1")
			fmt.Println("\nFlags")
			delCmd.PrintDefaults()
		}
		expenses := loadExpense()
		if len(expenses) == 0 {
			fmt.Println("There are no expenses yet")
		}

		deleteExpense(*id)
	case "monthlyExpense":
		monthlyCmd := flag.NewFlagSet("monthlyExpense", flag.ExitOnError)

		month := monthlyCmd.Int("month", 0, "Provide the months number")
		monthlyCmd.Parse(os.Args[2:])
		monthlyCmd.Usage = func() {
			fmt.Println("calculate  the expense for the month")
			fmt.Println("Usuage")
			fmt.Println("expense-tracker list")
			monthlyCmd.PrintDefaults()

		}

		sum := ExpensesByMonth(*month)
		fmt.Printf("Total expenses for month %d: $%.2f\n", *month, sum)
	case "help":
		fmt.Println("Expense Tracker CLI")
		fmt.Println("\nUsage:")
		fmt.Println("  expense-tracker <command> [flags]")
		fmt.Println("\nCommands:")
		fmt.Println("  add             Add a new expense")
		fmt.Println("  list            List all expenses")
		fmt.Println("  delete          Delete an expense")
		fmt.Println("  summary         Show total expenses")
		fmt.Println("  monthlyExpense  Show expenses for a month")
		fmt.Println("\nUse 'expense-tracker <command> -h' for command-specific help")

	}

}
