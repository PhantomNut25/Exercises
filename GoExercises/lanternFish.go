package main

import (
	"fmt"
)

func main() {
	//Initialize timers of lantern fishes
	fishes := make([]int, 0)
	fishes = append(fishes, 3, 4, 3, 1, 2)
	var days = 0

	for days < 18 {
		days++
		fmt.Print("After day ", days, " : ")

		for i := 0; i < len(fishes); i++ {
			if fishes[i] == 0 {
				fishes[i] = 7
				fishes = append(fishes, 9)

			}
			fishes[i] = fishes[i] - 1
		}

		fmt.Println(fishes)

	}

}
