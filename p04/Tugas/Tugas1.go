package main

import "fmt"

// Fungsi fact untuk menghitung faktorial dari suatu angka
func fact(n int) int {
    if n == 1 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    // Fungsi dibawah sama saja dengan
    // 7 x (6 x (5 x (4 x (3 x (2 x (1 x 1))))))
    fmt.Println(fact(7))
}
