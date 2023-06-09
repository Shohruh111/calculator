package main

import (
	"fmt"

	"calculator/calc"
)

func main() {
	fmt.Println(calc.Add(40, 50))
	fmt.Println(calc.Div(40, 50))
	fmt.Println(calc.Mult(40, 50))
	fmt.Println(calc.Sub(40, 50))

}
