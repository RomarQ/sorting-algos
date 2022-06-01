package main

import "fmt"

func main() {
	list := []int{5, 1, 4, 6, 3, 2}
	aux := make([]int, len(list))

	copy(aux, list)
	println("Quicksort (Right Most)")
	fmt.Println(aux)
	quickSort2(aux, partitionRightMost)
	fmt.Println(aux)

	copy(aux, list)
	println("Quicksort (Left Most)")
	fmt.Println(aux)
	quickSort2(aux, partitionLeftMost)
	fmt.Println(aux)

	copy(aux, list)
	println("Quicksort (Middle)")
	fmt.Println(aux)
	quickSort(aux, 0, len(aux)-1)
	fmt.Println(aux)
}

func quickSort(list []int, left int, right int) {
	if left >= 0 && left < right {
		p := partition(list, left, right)
		quickSort(list, left, p)
		quickSort(list, p+1, right)
	}
}

func partition(list []int, left int, right int) int {
	pivot := list[(right+left)/2]

	for {
		for list[left] < pivot {
			left += 1
		}
		for list[right] > pivot {
			right -= 1
		}

		if left >= right {
			return right
		}

		aux := list[right]
		list[right] = list[left]
		list[left] = aux
	}
}

func quickSort2(list []int, partition func(list []int) int) {
	if len(list) > 1 {
		p := partitionLeftMost(list)
		quickSort2(list[:p], partition)
		quickSort2(list[p:], partition)
	}
}

func partitionRightMost(list []int) int {
	pivotIndex := len(list) - 1
	pivot, i := list[pivotIndex], 0

	for idx, el := range list {
		if el < pivot {
			list[idx] = list[i]
			list[i] = el
			i += 1
		}
	}

	aux := list[i]
	list[i] = list[pivotIndex]
	list[pivotIndex] = aux

	return i
}

func partitionLeftMost(list []int) int {
	pivotIndex := 0
	pivot, i := list[pivotIndex], len(list)-1

	for j := i; j > pivotIndex; j-- {
		if list[j] > pivot {
			aux := list[j]
			list[j] = list[i]
			list[i] = aux
			i -= 1
		}
	}

	aux := list[i]
	list[i] = list[pivotIndex]
	list[pivotIndex] = aux

	return i + 1
}
