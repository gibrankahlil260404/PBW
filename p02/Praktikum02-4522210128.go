package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	daftarMahasiswa := make(map[string]Mahasiswa)

	daftarMahasiswa["123456789"] = Mahasiswa{"Rangga", "4522210118", "Pariwisata"}
	daftarMahasiswa["987654321"] = Mahasiswa{"Nabilah", "45222210129", "Ilmu Komunikasi"}

	fmt.Println("Daftar Mahasiswa:")
	for _, mhs := range daftarMahasiswa {
		fmt.Println("- ", mhs.Nama, mhs.NPM)
	}

	var namaBaru, npmBaru string
	fmt.Println("\nMasukkan nama mahasiswa baru:")
	fmt.Scanln(&namaBaru)

	fmt.Println("Masukkan NPM mahasiswa baru:")
	fmt.Scanln(&npmBaru)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan jurusan mahasiswa baru:")
	jurusanBaru, _ := reader.ReadString('\n')
	jurusanBaru = strings.TrimSpace(jurusanBaru)

	daftarMahasiswa[npmBaru] = Mahasiswa{namaBaru, npmBaru, jurusanBaru}

	fmt.Println("\nDaftar Mahasiswa Setelah Penambahan:")
	for _, mhs := range daftarMahasiswa {
		fmt.Println("- ", mhs.Nama, mhs.NPM)
	}

	var npmCari string
	fmt.Println("\nMasukkan NPM mahasiswa yang ingin dicari:")
	fmt.Scanln(&npmCari)

	mhs, found := daftarMahasiswa[npmCari]
	if found {
		fmt.Println("\nData Mahasiswa dengan NPM", npmCari, ":")
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("NPM:", mhs.NPM)
		fmt.Println("Jurusan:", mhs.Jurusan)
	} else {
		fmt.Println("\nMahasiswa dengan NPM", npmCari, "tidak ditemukan.")
	}
}
