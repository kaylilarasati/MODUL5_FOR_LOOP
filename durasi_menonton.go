package main
import "fmt"

func main() {
	var n, i int
	var jam, menit, sisa, durasi int

	fmt.Scan(&n)
	durasi = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&menit)
		durasi = durasi + menit

	}
	jam = durasi / 60
	sisa = durasi % 60
	menit = sisa
	fmt.Println(jam, "jam", menit, "menit")
}