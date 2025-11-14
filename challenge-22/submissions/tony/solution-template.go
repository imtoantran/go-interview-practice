package main

import (
	"fmt"
)

func main() {
	// Standard U.S. coin denominations in cents
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

// MinCoins returns the minimum number of coins needed to make the given amount.
// If the amount cannot be made with the given denominations, return -1.
func MinCoins(amount int, denominations []int) int {
	// TODO: Implement this function
	minCoin, _ := SolveProblem(amount, denominations)
	return minCoin
}

// CoinCombination returns a map with the specific combination of coins that gives
// the minimum number. The keys are coin denominations and values are the number of
// coins used for each denomination.
// If the amount cannot be made with the given denominations, return an empty map.
func CoinCombination(amount int, denominations []int) map[int]int {
	// TODO: Implement this function
	_, coinCombination := SolveProblem(amount, denominations)
	return coinCombination
}
func SolveProblem(amount int, denominations [] int) (int, map[int]int) {
	for i := len(denominations) - 1; i > -1; i-- {
		_result := make(map[int]int)
		_amount:=0
		_count:=0
		_current_coin_pos:= i
		for (_amount < amount) {
			if _current_coin_pos < 0 {
				break
			}
			current_coint:=denominations[_current_coin_pos]
			if current_coint > amount {break}
			if (current_coint + _amount > amount) {
				_current_coin_pos--
				continue
			}
			_amount+=current_coint
			_result[current_coint]+=1
			_count+=1
		}
		if _amount == amount {
			return _count, _result
		}
	}
	return -1, make(map[int]int)
}