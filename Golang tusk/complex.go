package main

import (
	"fmt"
)

func main(){
	num1 := complex(10,21)
	num2 := 10 + 5i
	num1Imagine := imag(num1)
	num1Real := real(num1)

	fmt.Printf("%v %v\n", num1,num2)	
	fmt.Printf("%v\t%v\n", num1Real,num1Imagine)	

}