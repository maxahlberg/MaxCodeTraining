package fizbuz

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func fiz(input int) (output int) {
	output = 42
	return output
}

func buz(input int) (output2 string) {
	output2 = strconv.Itoa(input)
	if math.Mod(float64(input), 3) == 0 {
		output2 = "Fizz"
	}
	if math.Mod(float64(input), 5) == 0 {
		output2 = "BuzZzzz"
	}
	if math.Mod(float64(input), 3) == 0 && math.Mod(float64(input), 5) == 0 {
		output2 = "FiiiIIiiiz BuuuuuZZzzzz"
	}

	return output2
}

//func correct(answer bool) {
//	if answer
//	fmt.Println("Fly you foool!")
//}

func Fizbuzmain() {
	fmt.Println("Hello Human, Welcome to Fiz BuZzzzz!")
	time.Sleep(4 * time.Second)
	fmt.Println("Enter a human number: ")
	var input int
	fmt.Scanln(&input)
	time.Sleep(1 * time.Second)
	fmt.Println("Your meaningless input was enterpreted as: ", input)

	toHuman := fiz(input)
	time.Sleep(4 * time.Second)
	fmt.Println("You should know that the meaning of life the universe and everything, is:", toHuman)

	toFizBuz := buz(input)
	time.Sleep(5 * time.Second)
	fmt.Println("The fiz buz output is:", toFizBuz)

}
