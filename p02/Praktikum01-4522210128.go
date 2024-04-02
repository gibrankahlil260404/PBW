package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Masukkan nama Anda: ")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')

	fmt.Print("Masukkan usia Anda: ")
	usiaInput, _ := reader.ReadString('\n')
	usiaInput = strings.TrimSpace(usiaInput) 
	usia, err := strconv.Atoi(usiaInput)
	if err != nil {
		fmt.Println("Input usia tidak valid.")
		return
	}

	var kategoriUsia string
	if usia < 18 {
		kategoriUsia = "anak-anak"
	} else {
		kategoriUsia = "dewasa"
	}

	fmt.Printf("Selamat datang %s, Anda termasuk kategori %s.", nama, kategoriUsia)
}
