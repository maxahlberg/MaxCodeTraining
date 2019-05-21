package main

import "fmt"

func fib(input int8) (resulting int8) {
	v0 := int8(0)
	v1 := int8(1)

	for i := int8(1); i < input; i++ {
		resulting = v0 + v1
		v0 = v1
		v1 = resulting
	}
	return resulting
}


func fibRec(input int8) int8{
	if input == 0 || input == 1{
	return input
	}
	return fibRec(input -2) + fibRec(input -1)

}

func main() {
	fmt.Println("Hello Max, input a number:  ")
	var input int8
	fmt.Scanln(&input)
	fmt.Println("The input was: ", input)

	var output int8
	var output2 int8
	output = fib(input)
	output2 = fibRec(input)

	fmt.Println("The fibonatchi number", input, "is", output)
	fmt.Println("According to the recursive method it is: ", output2)

}
