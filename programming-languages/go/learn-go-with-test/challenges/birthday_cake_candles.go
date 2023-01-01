package challenges

import "sort"

func birthdayCakeCandles(arr []int) int {
	
	// how to sort []int32
	// data := []int32{1, 3, 4, 2}
	// sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })

	sort.Ints(arr)
	var greatest, counter int 

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			counter++
		} else {
			counter = 0
		}

		if counter > greatest {
			greatest = counter
		}
	}

	return greatest+1
}

