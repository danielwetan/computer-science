package main

import (
	"fmt"
	"reflect"
)

func main() {
	got := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	want := 12.75
	// got := findMaxAverage([]int{0, 4, 0, 3, 2}, 1)
	// want := float64(4)
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("test failed, expected %v, but got %v\n", want, got)
		return
	}

	fmt.Println("test passed")
}

func findMaxAverage(nums []int, k int) float64 {
	var (
		start          = 0
		end            = k - 1
		result float64 = 0
		max    float64 = 0
	)

	for end < len(nums) {
		result = sumSubArray(nums[start:end+1]) / float64(k)
		if max == 0 {
			max = result
		}
		if result > max && result != 0 {
			max = result
		}

		start++
		end++
	}

	return max
}

func sumSubArray(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println(sum)

	return float64(sum)
}
