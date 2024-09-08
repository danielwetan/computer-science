package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	want := 49
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("test failed, expected %v, but got %v\n", want, got)
		return
	}

	fmt.Println("test passed")
}

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	max := 0

	for start < end {
		n := end - start
		sum := n * int(math.Min(float64(height[start]), float64(height[end])))

		if sum > max {
			max = sum
		}

		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}

	return max
}
