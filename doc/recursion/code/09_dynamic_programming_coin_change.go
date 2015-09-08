package main

import (
	"fmt"
	"math"
)

func getChange(amount int, coins []int) int {
	storage := []int{0}
	for a := 1; a <= amount; a++ {
		storage = append(storage, math.MaxInt32)
		for _, coin := range coins {
			if a >= coin {
				if storage[a] > 1+storage[a-coin] {
					// retrieve from storage
					storage[a] = 1 + storage[a-coin]
				}
			}
		}
	}
	return storage[amount]
}

// Find this minimum number of coins needed to make change fo x amount
func main() {
	coins := []int{1, 5, 7, 9, 11}

	fmt.Println(getChange(6, coins)) // 2
	//we need 2 coins(1 and 5) to make 6 cents

	fmt.Println(getChange(16, coins)) // 2
	//we need 2 coins(7 and 9) to make 16 cents

	fmt.Println(getChange(25, coins))  // 3
	fmt.Println(getChange(250, coins)) // 24
}
