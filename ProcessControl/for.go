package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1..100 is", sum)
	while()
	soooon()
	continued()
}

// Golang没有while关键字，所以只能改写for语句
func while() {
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}
}

// 死循环
func soooon() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writing inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

// continue
func continued() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum)
}

/*
Int63n表示返回一个小于n的非负整数
除了`rand.Int63n()`方法，在Go标准库中还有其他的方法可以生成随机数。以下是一些与`Int63n`功能类似的方法：

1. `rand.Intn(n int) int` ：生成0到n-1之间的随机整数。

2. `rand.Int() int` ：生成64位的随机整数。

3. `rand.Float64() float64` ：生成0.0到1.0之间的随机浮点数。

4. `rand.Float32() float32` ：生成0.0到1.0之间的随机浮点数。

这些方法都可以用来生成随机数，使用方式与`Int63n`类似。但是需要根据具体需求选择不同的方法以及对应的参数。
*/
