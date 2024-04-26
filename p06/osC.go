package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileInfo, err := os.Stat("GibranKahlil")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("GibranKahlil adalah sebuah direktori.")
	} else {
		fmt.Println("GibranKahlil adalah sebuah file.")
	}
}
