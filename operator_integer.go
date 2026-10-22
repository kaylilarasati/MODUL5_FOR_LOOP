package main

import "fmt"

func main() {
	var n, i, h int
	var b1, b2, b3, b4, b5, b6 int
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Scan(&b1, &b2, &b3, &b4, &b5, &b6)
		h = b1 + b2 - b3*b4/b5%b6
		fmt.Println(h)
	}
}
