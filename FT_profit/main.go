package main

import "fmt"

func Ft_profit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit

}

func main() {
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1})) // resultat : 0
}
