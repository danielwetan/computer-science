package challenges

import (
	"sort"
)

/*
NEED TO REFACTOR LATER
There are N children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

1. Each child must have at least one candy.
2. Children with a higher rating get more candies than their neighbors.

Input:
    A = [1, 5, 2, 1]

Output:
    7

Explanation:
    Candies given = [1, 3, 2, 1]
*/

func distributeCandy(arr []int) int {
	sort.Ints(arr)
	candies := len(arr)
	l := len(arr)

	for i := 0; i < len(arr); i++ {
		l--
		if i == len(arr)-1 {
			break
		}

		if arr[i] < arr[i+1] {
			candies = candies + l
		}
	}

	return candies
}

// func distributeCandy(arr []int) int {

// 	count := len(arr)

// 	nums := map[int]int{}

// 	for _, v := range arr {

// 		nv, found := nums[v]
// 		if found {
// 			nums[v] = nv + 1
// 		} else {
// 			nums[v] = 1
// 		}
// 	}

// 	for k, v := range nums {
// 		fmt.Println(k, v)
// 	}

// 	// fmt.Println(nums)

// 	// fmt.Println(len(nums))

// 	return count
// }
