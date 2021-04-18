package main

import "fmt"

func main() {
	var nums []int = []int{1, 1, 5, 5, 3, 4, 3, 1, 1, 5, 5, 5, 5, 3, 3, 4, 3, 4, 4, 1}
	highestRank(nums)
}

func highestRank(nums []int) map[int]int {
	var notes = map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

	for key := range notes {
		var b int = 0
		for j := 0; j < len(nums); j++ {
			if key == nums[j] {
				b++
			}
		}

		notes[key] = b
	}

	fmt.Println(notes)
	return notes
}
