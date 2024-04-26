package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	

	err = os.Chmod("GibranKahlil", 0128)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Println("Izin 'GibranKahlil' telah diubah menjadi 0128.")
}
