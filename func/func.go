package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)
}

/*
命名返回值则是在函数签名中声明了一个变量名，相当于将该变量作为返回值使用。
在第二段代码中，我们在函数签名中使用了 result int，
这意味着我们在函数内部可以直接使用 result 变量，
并在函数结束时将其返回，而不需要在 return 语句中显式声明返回值。
这使得代码更加清晰，易于阅读。

func sum(number1 string, number2 string) int {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    return int1 + int2
}
*/

func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}
