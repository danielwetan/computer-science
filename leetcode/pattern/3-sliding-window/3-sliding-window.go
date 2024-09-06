package main

import (
	"fmt"
	"reflect"
)

func main() {
	got := slidingWindow([]int{2, 1, 5, 1, 3, 2}, 3)
	want := 9

	fmt.Println(got)
	fmt.Println(want)
	if !reflect.DeepEqual(got, want) {
		fmt.Println("test failed")
	}
}

func slidingWindow(nums []int, length int) int {
	var (
		start   = 0
		end     = length - 1
		largest = 0
	)

	for end < len(nums) {
		sum := calculateSubArray(nums[start : end+1])
		if sum > largest {
			largest = sum
		}

		start++
		end++
	}

	return largest
}

func calculateSubArray(nums []int) int {
	result := 0

	for _, v := range nums {
		result += v
	}

	return result
}
