package main

import (
	"fmt"
	"os"
	"strconv"
)

var data = []string{"cloudacademy", "was", "here"}

func getElement(index int) string {
	return data[index]
}

func main() {
	var idx int
	defer fmt.Printf("defer called - idx is %d \n", idx)

	//note: 1st arg is always supplied by go runtime
	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	}

	arg := os.Args[1]
	fmt.Printf("argument passed in: %s\n", arg)
	idx, _ = strconv.Atoi(arg)
	s := getElement(idx)
	fmt.Printf("data at %d is: %s\n", idx, s)
}
