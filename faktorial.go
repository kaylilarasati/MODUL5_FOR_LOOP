package main
import "fmt"

func main() {
	var i, hasil int
	var a int
	fmt.Scan(&a)
	hasil = 1

	for i = 1;i <= a; i+=1{
		hasil = hasil * i
	}
	fmt.Print(hasil)
}

//pseudocode
// program faktorial
// kamus
// a, i, hasil : integer
// algoritma
// input(a)
// hasil = 1
// for i = 1 to a do
//     hasil = hasil * i
// end for
// output(hasil)