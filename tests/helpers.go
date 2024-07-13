package tests

import (
	"math/rand"
	"sort"
)

type comarableValues interface {
	~int | ~float32 | ~float64 | ~string
}

func getArrayOfRandomNonNegativeUniqueValues(size int) []int {
	seen_map := make(map[int]bool)
	var values []int
	for i := 0; i < size; i++ {
		ran_num := rand.Intn(size)
		_, seen := seen_map[ran_num]
		for seen {
			ran_num = rand.Intn(size)
			_, seen = seen_map[ran_num]
		}
		seen_map[ran_num] = true
		values = append(values, ran_num)
	}

	return values
}

func sortArrayAsc[T comarableValues](array []T) []T {
	//sorted := array
	var sorted []T
	for _, value := range array {
		sorted = append(sorted, value)
	}
	sortFunc := func(i, j int) bool {
		return sorted[i] < sorted[j]
	}
	sort.Slice(sorted, sortFunc)

	return sorted
}

func sortArrayDesc[T comarableValues](array []T) []T {
	//sorted := array
	var sorted []T
	for _, value := range array {
		sorted = append(sorted, value)
	}
	sortFunc := func(i, j int) bool {
		return sorted[i] > sorted[j]
	}
	sort.Slice(sorted, sortFunc)
	return sorted
}

func arraysEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func getIntArrayFromAnyArray(anyArr []any) []int {
	var intArr []int
	for _, value := range anyArr {
		intArr = append(intArr, value.(int))
	}
	return intArr
}
