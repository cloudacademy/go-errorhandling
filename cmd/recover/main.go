package main

import "fmt"

var data = []string{"cloudacademy", "was", "here"}

func systemStart() {
	fmt.Println("system starting...")

	defer func(msg string) {
		if r := recover(); r != nil {
			fmt.Println("RECOVERED!!")
		}
		fmt.Println(msg)
	}("blah")

	var nonexistent = data[3] //causes runtime panic!!
	fmt.Println(nonexistent)

	fmt.Println("system started")
}

func main() {
	systemStart()
}
