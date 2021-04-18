package main

import "fmt"

func main() {
	var listOfNotes []int = []int{1, 1, 5, 5, 3, 4, 3, 1, 1, 5, 5, 5, 5, 3, 3, 4, 3, 4, 4, 1}
	highestRank(listOfNotes)
}

func highestRank(listOfNotes []int) map[int]int {
	var notes = map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

	for key := range notes {
		var b int = 0
		for j := 0; j < len(listOfNotes); j++ {
			if key == listOfNotes[j] {
				b++
			}
		}

		notes[key] = b
	}
	for key, val := range notes {
		fmt.Printf("Оценка %d в количестве %d\n", key, val)
	}
	return notes
}
