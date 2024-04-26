package main

import (
	"fmt"
	"os"
)

func main() {
	// Membuat direktori baru.
	err := os.Mkdir("", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori '' berhasil dibuat.")
}
