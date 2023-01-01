package challenges

/*
Cari bilangan positif terkecil yg tidak ada dalam array

input: [4, 7, 0, 2, 1, 3, 5, 9, 10, 18, -1]
output: 6
*/

func smallestIntNotPresent(arr []int) int {

	smallestNum := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallestNum {
			smallestNum = arr[i]
		}
	}

	return next(smallestNum, arr)
}

func next(target int, arr []int) int {
	for _, num := range arr {
		if target == num {
			return next(target+1, arr)
		}
	}

	return target
}



// func smallestIntNotPresent(arr []int) int {
// 	arr = sorting(arr)

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i]+1 != arr[i+1] {
// 			return arr[i] + 1
// 		}
// 	}

// 	return 0
// }

// func sorting(arr []int) []int {
// 	noSwap := true

// 	for i := len(arr); i > 0; i-- {
// 		noSwap = true

// 		for j := 0; j < i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				temp := arr[j]
// 				arr[j] = arr[j+1]
// 				arr[j+1] = temp
// 				noSwap = false
// 			}
// 		}

// 		if noSwap { break }
// 	}

// 	return arr
// }
