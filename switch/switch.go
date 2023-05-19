package main

import (
	"fmt"
)

func main() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}

/*
请注意，由于 num 为 15（小于 50），因此它与第一个 case 匹配。
但是，num 不大于 100。 由于第一个 case 语句包含 fallthrough 关键字，
因此逻辑会立即转到下一个 case 语句，而不会对该 case 进行验证。
因此，在使用 fallthrough 关键字时必须谨慎。 该代码产生的行为可能不是你想要的。
*/
