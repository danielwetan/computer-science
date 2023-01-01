package challenges

// func minMaxSum(arr []int32) (int32, int32) {

// 	var min, max int32

// 	for i := 0; i < len(arr); i++ {
// 		if (i != 0) {
// 			max += arr[i]
// 		}
		
// 		if (i != len(arr)-1) {
// 			min += arr[i]
// 		}
		
// 	}

// 	return min, max 
// }


func minMaxSum(arr []int32) (int64, int64) {

	var min, max int64 = int64(arr[0]), 0
	var arrSum int64 = 0

	for i := 0; i < len(arr); i++ {
		digit := int64(arr[i])
		if min > digit {
			min = digit
		}
		
		if max < digit {
			max = digit
		}

		arrSum += digit
	}

	return arrSum-max, arrSum-min 
}