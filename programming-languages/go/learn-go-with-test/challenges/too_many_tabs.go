package challenges

import (
	"strings"
)

/*

Kelompokan domain yang sama menjadi satu tabs,
kemudian hitung jumlah tab yang dihasilkan

Contoh:
input: ["www.google.com", "www.google.com/id", "www.google.co.id", "www.google.co.id/account", "www.ruangguru.com/ruangbelajar", "roboguru.ruangguru.com", "roboguru.ruangguru.com/faq"]
output: 4

input: ["www.google.com", "www.google.com/id", "www.google.com/id/account", "www.google.com/id/account/reset", "www.google.com/id/account/forget_password"]
output: 1
*/

func tooManyTabs(arr []string) int {
	m := make(map[string]string)

	for _, domain := range arr {
		cleanDomain := strings.Split(domain, "/")[0]
		m[cleanDomain] = cleanDomain
	}

	return len(m)
}
