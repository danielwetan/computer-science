package challenges

/*
Cek apakah ada angka dari masing-masing array yang ketika dijumlahkan 
nilainya sama dengan target yang di set

Input: 
a: [1, 2, 3]
b: [10, 20, 30, 40]
v: 42 // 40 + 2

Output:
true
*/

func sumOfTwo(a, b []int, v int) bool {
	remaining := map[int]int{}

	for _, n := range a {
		r := v - n
		remaining[r] = r
	}

	for _, n := range b {
		_, found := remaining[n]
		if found {
			return true
		}
	}

	return false
}
