package main
import "fmt"

func main() {
	var a, b int
	var i int
	fmt.Scan(&a, &b)
	for i = a; i <= b; i+=1{
		fmt.Print(i, " ")
	}

}




//pseudocode
// program baris_bilangan
// kamus
// a, b, i : integer
// algoritma
// input(a, b)
// for i = a to b do
//     output(i)
// end for