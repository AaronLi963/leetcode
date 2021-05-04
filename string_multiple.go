package main

import (
	"fmt"
)

func strProduct(strA, strB string) string {
	m := len(strA)
	n := len(strB)

	ans := make([]byte, m+n)
	for i := range ans {
		ans[i] = '0'
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			prod := (strA[i] - '0') * (strB[j] - '0')
			ans[i+j+1] += prod % 10
			ans[i+j] += prod / 10
		}
	}

	for i := m + n - 1; i >= 0; i-- {
		if ans[i] > '9' {
			diff := ans[i] - '0'
			ans[i] = diff%10 + '0'
			ans[i-1] += diff / 10
		}
	}
	for i := range ans {
		if ans[i] != '0' {
			return string(ans[i:])
		}
	}
	return "0"
}

func main() {
	A := "12"
	B := "34"
	prod := strProduct(A, B)
	fmt.Println(prod)
}
