package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Not enough number of arguements")
		fmt.Println("Use 'expense-tracker help' to see available commands")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		handleAddCmd()

	case "list":
		handleListCmd()

	case "summary":
		handleSumCmd()

	case "delete":
		handleDelCmd()

	case "monthlyExpense":
		handleMonthCmd()
	case "update":
		handleUpdateCmd()
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
func printCommandUsage(cmd *flag.FlagSet, usageText string) {
	fmt.Println(usageText)
	fmt.Println("\nFlags:")
	cmd.PrintDefaults()
}
func handleAddCmd() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCmd.String("description", "", "add your new expense")
	amount := addCmd.Float64("amount", 0, "the expense amount")
	dateStr := addCmd.String("date", "", "Optional date (YYYY-MM-DD)")
	var date time.Time
	var err error
	addCmd.Parse(os.Args[2:])

	if *dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", *dateStr)
		if err != nil {
			fmt.Println("Invalid date format. Use YYYY-MM-DD")
			return
		}
	}

	if *description == "" || *amount <= 0 {
		printCommandUsage(addCmd, "Add a new expense\nUsage: expense-tracker add -description=\"Lunch\" -amount=50 [-date=2026-04-08]")
		return
	}
	exp, err := AddExpense(*description, *amount, date)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Added expense: ID=%d\n", exp.ID)

}
func handleListCmd() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.Parse(os.Args[2:])
	printCommandUsage(listCmd, "expense-tracker list")
	expenses, err := ListExpenses()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)

	}
	if len(expenses) == 0 {
		fmt.Println("There are no expenses yet")
	}
	for _, r := range expenses {
		fmt.Printf("#%d\t%s\t%s\t%.2f\n",
			r.ID,
			r.Date.Format("2006-01-02"),
			r.Description, r.Amount,
		)
	}
}
func handleSumCmd() {
	sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	sumCmd.Parse(os.Args[2:])
	printCommandUsage(sumCmd, "expense-tracker sumarry")
	sum, err := TotalExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total expenses: $%.2f\n", sum)

}
func handleDelCmd() {
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := delCmd.Int("id", 0, "the id of the expense you want to delete")
	delCmd.Parse(os.Args[2:])
	printCommandUsage(delCmd, " expense-tracker delete -id=1")
	expenses, err := loadExpense()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if len(expenses) == 0 {
		fmt.Println("There are no expenses yet")
	}

	err = deleteExpense(*id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Expense deleted successfully (ID: %d)\n", *id)
}
func handleMonthCmd() {
	monthlyCmd := flag.NewFlagSet("monthlyExpense", flag.ExitOnError)

	month := monthlyCmd.Int("month", 0, "Provide the months number")
	monthlyCmd.Parse(os.Args[2:])
	printCommandUsage(monthlyCmd, "expense-tracker list")

	sum, err := ExpensesByMonth(*month)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total expenses for month %d: $%.2f\n", *month, sum)
}
func handleUpdateCmd() {
	upCmd := flag.NewFlagSet("updating amount of an expense", flag.ExitOnError)
	description := upCmd.String("desc", "", "describe  the expense to be updated")
	amount := upCmd.Float64("amount", 0, "amount  to be updated")
	upCmd.Parse(os.Args[2:])
	exp, err := updateExpense(*description, *amount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Expense updated successfully (ID: %d)\n", exp.ID)

}
