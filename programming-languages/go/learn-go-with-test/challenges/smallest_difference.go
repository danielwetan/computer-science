package challenges

import (
	"math"
	"sort"
)

func smallestDifference(arr []int) int {
	sort.Ints(arr)

	diff := math.MaxInt64

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < diff {
			diff = arr[i+1] - arr[i]
		}
	}

	return diff
}
