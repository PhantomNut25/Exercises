package main

import "fmt"

func main() {
	//Constant array length
	const N = 10
	numbers := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println("product: ", strangeProduct(numbers))
	evenOdd(numbers)
	fmt.Println("Minimum odd: ", minOdd(numbers))
}

// Multiply odd numbers in the sequence
func strangeProduct(numbers []int) int {
	var prod = 1
	for i := 0; i < len(numbers); i++ {
		if (numbers[i] > 7) && (numbers[i]%4 == 0) {
			prod *= numbers[i]
		}
	}
	return prod
}

// Find odd and even numbers in the sequence
func evenOdd(numbers []int) {
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, " pari")
		} else {
			fmt.Println(number, " dispari")
		}
	}
}

// Find the minimum odd number in the sequence
func minOdd(numbers []int) int {
	var min int
	var index int

	for i, number := range numbers {
		if number%2 != 0 {
			min = number
			index = i
			break
		}
	}

	for j := index; j < len(numbers)-index; j++ {
		if numbers[j]%2 != 0 {
			if numbers[j] < min {
				min = numbers[j]
			}
		}
	}

	return min
}
