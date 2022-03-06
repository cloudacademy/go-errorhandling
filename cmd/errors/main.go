package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func createDir() error {
	_, err := ioutil.TempDir("", "cloudacademy")
	if err != nil {
		return fmt.Errorf("problem encountered creating dir: %v", err)
	}

	return nil
}

func divide(num1 float32, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0.0, errors.New("cannot divide by 0")
	} else {
		return num1 / num2, nil
	}
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("radius should be positive value")
	} else {
		return math.Pi * math.Pow(float64(radius), 2), nil
	}
}

func main() {
	err := createDir()
	if err != nil {
		log.Println(err)
	}

	result, err := divide(23.5, 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("answer %.2f\n", result)
	}

	if area, err := circleArea(3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("answer %.2f\n", area)
	}
}
