package challenges

/*
Pindahkan seluruh angka nol dari array
ke bagian paling kanan dari array tersebut.

Contoh:
Input: [0, 1, 0, 3, 12]
Output: [1, 3, 12, 0, 0]
*/

func moveZeroes(nums []int) []int  {
	var zeros, other, result []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros = append(zeros, 0)
		} else {
			other = append(other, nums[i])
		}
	}

	result = append(other, zeros...)
	return result
}

// func moveZeroes(arr []int) []int {
// 	result := arr
// 	zeroCount := 0

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == int(0) {
// 			zeroCount++
// 			// removeSlice(result, i)
// 		}
// 	}

// 	fmt.Println(zeroCount)

// 	zeros := make([]int, zeroCount)
// 	result = append(result[:len(result)-zeroCount], zeros...)

// 	return result
// }

// func removeSlice(arr []int, index int) []int {
// 	result := append(arr[:index], arr[index+1:]...)
// 	return result
// }
