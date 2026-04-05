# Clifunctions-ExpenseTracker
* https://roadmap.sh/projects/expense-tracker
# Expense Tracker CLI

A simple **Go command-line application** to track personal expenses.  
You can add, list, delete expenses, get a summary, or check monthly expenses.

---

## Features

- Add new expenses with a description and amount  
- List all recorded expenses  
- Delete expenses by ID  
- Show total expenses (summary)  
- Show expenses for a specific month  
- Per-command help with usage  

---

## Installation

1. **Clone the repository:**

```bash
git clone <https://github.com/joelyakubu269/Clifunctions-ExpenseTracker>
cd expense-tracker
````

2. **Build the binary (if not built yet or after code changes):**

```bash
go build -o expense-tracker
```

This will create an `expense-tracker` binary in the project directory.

3. **Optional:** Move the binary to your PATH to run it from anywhere:

```bash
mv expense-tracker ~/go/bin/
export PATH=$PATH:~/go/bin
```

---

## Usage

Run commands as:

```bash
./expense-tracker <command> [flags]
```

Use `./expense-tracker help` to see all available commands.

---

## Commands

### 1. Add a New Expense

```bash
./expense-tracker add -description "Lunch" -amount 12.50
```

* Adds a new expense with a description and amount.
* Both `-description` and `-amount` are **required**.

---

### 2. List All Expenses

```bash
./expense-tracker list
```

* Lists all recorded expenses with IDs, descriptions, and amounts.
* If no expenses exist, a message will indicate this.

---

### 3. Delete an Expense

```bash
./expense-tracker delete -id 3
```

* Deletes the expense with the specified ID.
* `-id` is **required**.

---

### 4. Show Total Expenses (Summary)

```bash
./expense-tracker summary
```

* Calculates and displays the **total amount of all expenses**.

---

### 5. Show Monthly Expenses

```bash
./expense-tracker monthlyExpense -month 4
```

* Shows total expenses for the specified month (1 = January, 2 = February, etc.).
* `-month` is **required**.

---

### 6. Help

```bash
./expense-tracker help
```

* Shows all commands and usage information.
* You can also run `-h` for command-specific help:

```bash
./expense-tracker add -h
./expense-tracker delete -h
./expense-tracker summary -h
```

---

## Notes

* Flags **must start with `-`**.
* Multi-word descriptions must be quoted: `"Grocery shopping"`
* Commands that require IDs or amounts will show an error if flags are missing.

```

---


