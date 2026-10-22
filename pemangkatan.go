package main
import "fmt"

func main(){
	var  i, hasil int
	var a, b int
	fmt.Scan(&a, &b)
	hasil = 1
	for i =0; i < b; i++ {
		hasil = hasil * a
	}
	fmt.Print(hasil)
}