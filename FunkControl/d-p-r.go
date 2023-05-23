package main

import (
	"fmt"
	"io"
	"os"
)

// 延迟调用输出的顺序相反（后进先出）
func main() {
	// deferf()

	/*
		highlow(2, 0)
		fmt.Println("Program finished successfully!")
	*/
	recoverd()
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

func deferf() {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newfile.Close()

	// 写入字符的代码实际上可以放到if语句外面，此处是为了程序的易读性
	if _, error = io.WriteString(newfile, "Learning Go!"+"\n"+
		"Horon_Lee"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

	newfile.Sync()
}

func recoverd() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
