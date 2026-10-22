package main

import "fmt"

func main() {
	var i, n int
	var tabungan_awal, kenaikan, total, penambahan int

	fmt.Scan(&n, &tabungan_awal, &kenaikan)

	total = tabungan_awal
	penambahan = kenaikan

	for i = 1; i <= n; i++ {
		penambahan = penambahan + kenaikan
		total = total + penambahan
	}

	fmt.Println(total)
}