package main

import (
	"fmt"
	"log"
	"os"
)

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return nil
	} else {
		return file
	}
}

func closeFile(file *os.File) {
	fmt.Println("closeFile called...")
	if file != nil {
		err := file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error! %v\n", err)
			os.Exit(1)
		}
	}
}

func func1() {
	defer fmt.Println("func1")
	func2()
}

func func2() {
	defer fmt.Println("func2")
}

func main() {
	file := createFile("cloudacademy.txt")
	defer closeFile(file)

	fmt.Fprintln(file, "defer stmt will always be called...")

	func1()
}
