package main

import "fmt"

func main() {
    // Mendefinisikan batas nilai dalam sebuah map
    batasNilai := map[string]int{
        "A": 90,
        "B": 80,
        "C": 70,
        "D": 60,
    }

    // Closure untuk menentukan nilai huruf
    getNilaiHuruf := func(nilaiUjian int) string {
        for huruf, batas := range batasNilai {
            if nilaiUjian >= batas {
                return huruf
            }
        }
        return "E"
    }

    // Menentukan nilai huruf untuk beberapa nilai ujian
    nilaiUjian := []int{85, 75, 95, 55, 65}
    for _, nilai := range nilaiUjian {
        fmt.Println("Nilai ujian", nilai, ":", getNilaiHuruf(nilai))
    }
}
