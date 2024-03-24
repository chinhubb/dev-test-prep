package main

import (
	"fmt"
)

// collection_1, collection_3 already sorted from min(0) to max
// collection_2 already sorted from max to min(0)
func merge(collection_1, collection_2, collection_3 []int) []int {
	totalLength := len(collection_1) + len(collection_2) + len(collection_3)
	result := make([]int, totalLength)
	i := 0
	j := 0
	k := 0

	for p := 0; p < totalLength; p++ {
		if i >= len(collection_1) {
			if j >= len(collection_2) {
				result[p] = collection_3[k]
				k++
			} else {
				result[p] = collection_2[j]
				j++
			}
			continue
		} else if j >= len(collection_2) {
			if k >= len(collection_3) {
				result[p] = collection_1[i]
				i++
			} else if collection_1[i] <= collection_3[k] {
				result[p] = collection_1[i]
				i++
			} else {
				result[p] = collection_3[k]
				k++
			}
			continue
		} else if k >= len(collection_3) {
			if collection_1[i] <= collection_2[j] {
				result[p] = collection_1[i]
				i++
			} else {
				result[p] = collection_2[j]
				j++
			}
			continue
		}

		if collection_1[i] <= collection_3[k] && collection_1[i] <= collection_2[j] {
			result[p] = collection_1[i]
			i++
		} else if collection_3[k] <= collection_1[i] && collection_3[k] <= collection_2[j] {
			result[p] = collection_3[k]
			k++
		} else {
			result[p] = collection_2[j]
			j++
		}
	}

	return result
}

func main() {
	collection_1 := []int{2, 5, 9}
	collection_2 := []int{20, 13, 7}
	collection_3 := []int{3, 6, 14}

	result := merge(collection_1, collection_2, collection_3)
	fmt.Println(result)
}
