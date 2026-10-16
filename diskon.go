	package main

	import "fmt"

	func main() {
		var harga, total1, total2 float64

		fmt.Scan(&harga)

		total1 = harga * 0.9
		total2 = total1 * 0.8

		fmt.Printf("%.2f\n", total2)
	}