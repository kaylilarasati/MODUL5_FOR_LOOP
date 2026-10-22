package main
import "fmt"

func main() {
    var n, i, prioritas int

    fmt.Scan(&n)

    for i = 1; i <= n && i <= 1000; i++ {
        prioritas = n - i + 1
        fmt.Println("Request", i, "diproses dengan prioritas", prioritas)
    }
}
