package main

import "fmt"

func Ft_coin(coins []int, amount int) int {
	numCoins := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			numCoins += amount / coins[i]
			amount %= coins[i]
		}
	}
	if amount == 0 {
		return numCoins
	}
	return -1
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 (5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
}
