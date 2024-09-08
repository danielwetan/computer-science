package main

import (
	"fmt"
	"reflect"
)

func main() {
	got := prefixSum([]int{1, 2, 3, 4, 5, 6}, 1, 3)
	want := 9
	if !reflect.DeepEqual(got, want) {
		fmt.Println("test failed")
		fmt.Printf("expected %d, but got %d\n", want, got)
	} else {
		fmt.Println("test passed")
	}
}

func prefixSum(nums []int, start int, end int) int {
	newArr := nums[start : end+1]
	result := 0

	for _, v := range newArr {
		result += v
	}

	return result
}
