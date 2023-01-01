package challenges

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	var i int32
	var space string = " "
	var hash string = "#"

	for i = 1; i < n+1; i++ {
		a := strings.Repeat(space, int(n-i))
		b := strings.Repeat(hash, int(i))
		fmt.Println(a + b)
	}
}
