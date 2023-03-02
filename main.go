package main

import (
	"fmt"
	calc "github.com/adheteguhdev/go-simple-calculator"
)

func main(){
	a := 10
	b := 5
	resultAdd := calc.Add(a, b)
	fmt.Printf("The result of %d + %d = ",a,b)
	fmt.Println(resultAdd)

	resultMultiply := calc.Multiply(a,b)
	fmt.Printf("The result of %d * %d = ",a,b)
	fmt.Println(resultMultiply)

	resultSubs := calc.Subtract(a,b)
	fmt.Printf("The result of %d - %d = ",a,b)
	fmt.Println(resultSubs)

	resultDivide := calc.Divide(a, b)
	fmt.Printf("The result of %d / %d = ", a,b)
	fmt.Println(resultDivide)
}