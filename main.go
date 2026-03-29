package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main () {
	fmt.Println("Type: [expense-tracker list] to know how to use the cli")
	if len(os.Args) < 2 {
		fmt.Println("Not enough number of arguements")
		return
	}
	command:= os.Args[1]
	switch command {
	case "add":
		if len(os.Args)< 6 {
			fmt.Println("not enough number of arguements")
			return
		}
		//description:= os.Args[3]
		description:= strings.Join(os.Args[3:4]," ")
		amount:= os.Args[5]
		val:=  ID(amount)
		AddExpense(description ,val)
	case "list":
		if len(os.Args) < 2 {
			fmt.Println("Not enough number of arguements")
			return
		}
		ListExpenses()
	case ""
	}
}
func ID(amount string) int {
	n,err:= strconv.Atoi(amount)
	if err!= nil {
		fmt.Println("Error:",err)
		fmt.Println("Input a valid number")
	}
	return n
}