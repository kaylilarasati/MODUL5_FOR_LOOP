package main
import "fmt"

func main(){
	type buku struct {
		judul string
		genre string
		tahunrilis string
		rating float64
	}
	var a buku
	var b buku
	fmt.Scan(&a.judul, &b.judul, &a.genre, &b.genre, &a.tahunrilis, &b.tahunrilis, &a.rating, &b.rating)
	fmt.Println("--- Data Buku ---")
	fmt.Println("Buku 1:")
	fmt.Println("Judul:", a.judul)
	fmt.Println("Genre:", a.genre)
	fmt.Println("Tahun Rilis:", a.tahunrilis)
	fmt.Println("Rating:", a.rating)

	fmt.Println("Buku 2:")
	fmt.Println("Judul:", b.judul)
	fmt.Println("Genre:", b.genre)
	fmt.Println("Tahun Rilis:", b.tahunrilis)
	fmt.Println("Rating:", b.rating)
}

// pseucode
// program buku
// kamus 
//     type buku <
//         judul : string
//         genre : string
//         tahunrilis : string
//         rating : real 
// 	 >
    
//     buku1, buku2 : buku

// Algoritma
//     input (buku1.judul, buku2.judul, buku1.genre, buku2.genre, buku1.tahunrilis, buku2.tahunrilis, buku1.rating, buku2.rating)
//     output ("--- Data Buku ---"
// 			 "Buku 1:"
// 			 "Judul:", buku1.judul
// 			 "Genre:", buku1.genre
// 			 "Tahun Rilis:", buku1.tahunrilis
// 			 "Rating:", buku1.rating
// 			 "Buku 2:"
// 			 "Judul:", buku2.judul
// 			 "Genre:", buku2.genre
// 			 "Tahun Rilis:", buku2.tahunrilis
// 			 "Rating:", buku2.rating
// 			 )
// endprogram
