package main
import "fmt"

func main(){
	var n, i, ganjil int

	fmt.Scan(&n)
	ganjil = 1
	for i = 1; i <= n && ganjil <= 1000; i++{
		fmt.Println(ganjil)
		ganjil += 2
	}
}
