	package main
	
	import (
	    "fmt"
	    "progo/ratarata" // Mengimpor package
	)
	
	func main() {
	    // Data nilai siswa
	    nilaiSiswa := []int{80, 75, 90, 85, 60}
	
	    // Menghitung rata-rata dengan fungsi dari package
	    rataRata := ratarata.HitungRataRata(nilaiSiswa)
	
	    // Menampilkan hasil
	    fmt.Println("Rata-rata nilai siswa:", rataRata)
	}
