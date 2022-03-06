package main

import (
	"errors"
	"fmt"
	"os"
)

func openFile(filename string) error {
	if _, err := os.Open(filename); err != nil {
		return fmt.Errorf("problem encountered %s: %w", filename, err)
	}

	return nil
}

func main() {
	err := openFile("cloudacademy.log")

	if err != nil {
		fmt.Printf("error detected: %s \n", err.Error())
		unwrappedErr := errors.Unwrap(err)
		fmt.Printf("orig error: %v\n", unwrappedErr)
	}
}
