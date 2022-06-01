package main

import "fmt"

func main() {
	list := []int{5, 1, 4, 6, 3, 2}
	fmt.Println(list)
	fmt.Println(bubbleSort(list))
}

func bubbleSort(list []int) []int {
	sortedList := make([]int, len(list))
	// Clone list
	copy(sortedList, list)

	// Sort list
	for range sortedList {
		for idx, el := range sortedList[1:] {
			if sortedList[idx] > el {
				sortedList[idx+1] = sortedList[idx]
				sortedList[idx] = el
			}
		}
	}

	return sortedList
}
