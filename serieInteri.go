package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}

	fmt.Println("prodotto: ", stranoProdotto(numeri))
	pariDispari(numeri)
	fmt.Println("minDispari: ", minDispari(numeri))
}

func stranoProdotto(numeri []int) int {
	var prod = 1
	for i := 0; i < len(numeri); i++ {
		if (numeri[i] > 7) && (numeri[i]%4 == 0) {
			prod *= numeri[i]
		}
	}
	return prod
}

func pariDispari(numeri []int) {
	for _, numero := range numeri {
		if numero%2 == 0 {
			fmt.Println(numero, " pari")
		} else {
			fmt.Println(numero, " dispari")
		}
	}
}

func minDispari(numeri []int) int {
	var min int
	var index int

	for i, numero := range numeri {
		if numero%2 != 0 {
			min = numero
			index = i
			break
		}
	}

	for j := index; j < len(numeri)-index; j++ {
		if numeri[j]%2 != 0 {
			if numeri[j] < min {
				min = numeri[j]
			}
		}
	}

	return min
}
