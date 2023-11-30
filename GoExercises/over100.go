package main

import "fmt"

func main() {
	var number int
	//Scan a sequence of numbers until a number is greater than 100, break with -1
	for {
		fmt.Scan(&number)
		if number == -1 {
			break
		}

		if number > 100 {
			fmt.Println(number)
			return
		}
	}
	fmt.Println("No number greater than 100")
}
