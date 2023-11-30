package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing inital balance")
	}
	//Read initial balance from main argument
	balance, _ := strconv.Atoi(os.Args[1])
	var expenses int
	//While balance is greater than 0 keep reading expences
	for balance > 0 {
		fmt.Scan(&expenses)

		balance -= expenses

		fmt.Println("balance: ", balance)

	}
	fmt.Println("Final balance: ", balance)
}
