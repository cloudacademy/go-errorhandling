package main

import (
	"errors"
	"fmt"
)

// ===================================

var divideByZeroError = errors.New("divide by zero not permitted")

func divide1(num1 float32, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0, divideByZeroError
	}

	result := num1 / num2
	return result, nil
}

func calc1(num1, num2 float32) {
	result, err := divide1(num1, num2)

	//errors.Is example
	if err != nil {
		switch {
		case errors.Is(err, divideByZeroError):
			fmt.Println("divide by zero detected")
		default:
			fmt.Printf("some other error: %s\n", err)
		}
	} else {
		fmt.Printf("result %.2f\n", result)
	}

	fmt.Println()
}

// ===================================

type MathError struct {
	LeftOperand  float32
	RightOperand float32
	Message      string
}

func (e *MathError) Error() string {
	return e.Message
}

func divide2(num1 float32, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0, &MathError{
			LeftOperand:  num1,
			RightOperand: num2,
			Message:      "divide by zero attempted",
		}
	}

	result := num1 / num2
	return result, nil
}

func calc2(num1, num2 float32) {
	result, err := divide2(num1, num2)

	//errors.As example
	if err != nil {
		var mathError *MathError
		switch {
		case errors.As(err, &mathError):
			fmt.Println("math error detected")
			fmt.Println(mathError.Message)
		default:
			fmt.Printf("some other error: %s\n", err)
		}
	} else {
		fmt.Printf("result %.2f\n", result)
	}

	fmt.Println()
}

// ===================================

func main() {
	fmt.Println("errors.Is example...")
	calc1(23.5, 0)

	fmt.Println("errors.As example...")
	calc2(23.5, 0)
}
