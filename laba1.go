package main

import (
	"fmt"
	"laba/bigdigit"
)

const (
	path1      = "digit1.txt"
	path2      = "digit2.txt"
	pathResult = "result.txt"
)

func main() {
	d1, _ := bigdigit.SetFile("bigdigit/digit.txt")
	d2, _ := bigdigit.SetFile("bigdigit/digit2.txt")
	fmt.Println(bigdigit.CmpDigit(d1, d2))
}
