package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError  = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "Gibran" {
		return NotFoundError
	}
	// sukses
	return nil
}

func main() {
	err := GetById("Gibran")
	if err != nil {
		switch {
		case errors.Is(err, ValidationError):
			fmt.Println("validation error")
		case errors.Is(err, NotFoundError):
			fmt.Println("not found error")
		default:
			fmt.Println("unknown error")
		}
	}
}
