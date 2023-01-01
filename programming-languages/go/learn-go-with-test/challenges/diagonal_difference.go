package challenges

func diagonalDifference(arr [][]int) int {
	var left, right, counter int

	for i := 0; i < len(arr); i++ {
		left += arr[i][counter]
		right += arr[i][len(arr)-1-counter]
		counter++
	}

	if left > right {
		return left-right
	}

	return right - left
}

