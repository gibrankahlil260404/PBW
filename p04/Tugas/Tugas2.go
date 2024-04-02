package main

import "fmt"

// Mengurutkan array dalam satu kali perulangan
func main() {
    fmt.Println("Bubble Sort")

    // Inisialisasi array dengan nilai
    var arrayNumber = [20]int{2, 1, 5, 3, 6, 4, 7, 8, 9, 10, 12, 11, 15, 16, 14, 13, 17, 18, 19, 20}

    // Variabel untuk menunjukkan apakah terjadi pertukaran dalam iterasi saat ini
    swapped := true

    for swapped {
        swapped = false
        for i := 0; i < len(arrayNumber)-1; i++ {
            if arrayNumber[i] > arrayNumber[i+1] {
                // Jika nilai pada indeks i lebih besar dari nilai pada indeks i+1, tukar posisinya
                arrayNumber[i], arrayNumber[i+1] = arrayNumber[i+1], arrayNumber[i]
                swapped = true
            }
        }
    }

    fmt.Println("Setelah dilakukan pengurutan")
    fmt.Println(arrayNumber)
}
