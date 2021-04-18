package main

import "fmt"

func main() {
	var nums []int = []int{1, 1, 5, 5, 3, 4, 3, 1, 1, 5, 5, 5, 5, 3, 3, 4, 3, 4, 4, 1}
	highestRank(nums)
}

func highestRank(nums []int) map[int]int {
	num := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var b int = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				b++
			}
		}
		_, ok := num[nums[i]]
		if !ok {
			num[nums[i]] = b

		}

	}
	printMostResult(num)
	printWorstResult(num)
	fmt.Println(num)
	return num
}
func printMostResult(num map[int]int) {
	var maxRank int
	var bestNote int
	for key, val := range num {
		if val > maxRank {
			maxRank = val
			bestNote = key
		}
	}
	fmt.Printf("Оценка %d получена %d раз(а)\n", bestNote, maxRank)
}
func printWorstResult(num map[int]int) {
	var maxRank int = 1000
	var worstNote int
	for key, val := range num {
		if val < maxRank {
			maxRank = val
			worstNote = key
		}
	}
	fmt.Printf("Оценка %d получена %d раз(а)\n", worstNote, maxRank)
}
