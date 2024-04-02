package main

import (
    "fmt"
    "math"
)

func main() {
    // Deklarasi jari-jari
    jariJari := 5.0

    // Menghitung luas dengan anonymous function
    luas := func(r float64) float64 {
        return math.Pi * r * r
    }(jariJari)

    // Menghitung keliling dengan anonymous function
    keliling := func(r float64) float64 {
        return 2 * math.Pi * r
    }(jariJari)

    // Menampilkan hasil
    fmt.Println("Luas lingkaran:", luas)
    fmt.Println("Keliling lingkaran:", keliling)
}
