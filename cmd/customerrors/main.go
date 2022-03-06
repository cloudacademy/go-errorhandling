package main

import (
	"fmt"
	"log"
	"math"
)

type NegativeRadiusError struct {
	Radius float64
}

func (e *NegativeRadiusError) Error() string {
	return fmt.Sprintf("radius is negative: %v", e.Radius)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &NegativeRadiusError{Radius: radius}
	} else {
		return math.Pi * math.Pow(float64(radius), 2), nil
	}
}

func main() {
	if area1, err := circleArea(-3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("answer %.2f\n", area1)
	}

	area2, err := circleArea(-3)

	if err != nil {
		switch e := err.(type) {
		case *NegativeRadiusError:
			log.Println("-ve radius detected!")
		default:
			log.Println(e)
		}
	} else {
		fmt.Printf("answer %.2f\n", area2)
	}
}
