package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	// if x is a negative number, sqrt() will return an error
	sqResult, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqResult)
	}
}

// funcKeyWord funcName (paramName paramType, paramName paramType) returnType { code }
func sum(x int, y int) int {
	return x + y
}

// function that has multiple return statements
// func funcName (paramName paramType, paramName paramType) (returnType, returnType) { code }
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
