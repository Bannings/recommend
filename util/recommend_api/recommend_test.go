package recommend_api

import (
	"testing"
)

func TestTopRecommendApplyV2(t *testing.T) {
	a := []string{"123", "456", "789", "234", "567", "891", "147", "258", "369"}
	b := Random(a)
	t.Log(b)
}

func maxProfit(prices []int) int {
	var profit int
	for i := 0; i <= len(prices)-2; i++ {
		for j := i + 1; j <= len(prices)-1; j++ {
			temp := prices[j] - prices[i]
			if temp > profit {
				profit = temp
			}
		}
	}
	return profit
}

func maxProfit1(prices []int) int {
	var profit int
	for i := 0; i <= len(prices)-2; i++ {
		for j := i + 1; j <= len(prices)-1; j++ {
			temp := prices[j] - prices[i]
			if temp > 0 {
				break
			}
		}
	}
	return profit
}
