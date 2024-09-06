package main

import (
	"fmt"
	"reflect"
)

func main() {
	got := twoPointers([]int{1, 2, 3, 4, 6}, 6)
	want := []int{1, 3}

	fmt.Println(got)
	fmt.Println(want)
	if !reflect.DeepEqual(got, want) {
		fmt.Println("test failed")
	}
}

func twoPointers(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
