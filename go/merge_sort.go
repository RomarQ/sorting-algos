package main

import "fmt"

func main() {
	list := []int{5, 1, 4, 6, 3, 2}
	fmt.Println(list)
	sortedList := mergeSort(list)
	fmt.Println(sortedList)
}

func mergeSort(list []int) []int {
	if length := len(list); length > 1 {
		middle := length / 2
		return merge(mergeSort(list[0:middle]), mergeSort(list[middle:length]))
	}

	return list
}

func merge(left []int, right []int) []int {
	sortedList := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			sortedList = append(sortedList, right[0])
			right = right[1:]
		} else {
			sortedList = append(sortedList, left[0])
			left = left[1:]
		}
	}

	return append(sortedList, append(left, right...)...)
}
