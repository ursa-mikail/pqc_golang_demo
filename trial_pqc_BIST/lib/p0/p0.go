package p0

import (
	"fmt"
)

var Name string = "DX-50017"
var print = fmt.Println

func Xello() string {
	return "ursa_00"
}

func Xello_() {
	print("ursa_00 00")
}

func UseFunc(f func(int, int) int,
	x, y int) {
	print("Ans: ", (f(x, y)))
}

func SumVals(x, y int) int {
	return (x + y)
}
