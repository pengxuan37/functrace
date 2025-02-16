package main

import (
	"fmt"

	"github.com/pengxuan37/functrace"
)

func A1() {
	defer functrace.Trace()()
	B1()
}

func B1() {
	defer functrace.Trace()()
	C1()
}

func C1() {
	defer functrace.Trace()()
	D()
}

func D() {
	defer functrace.Trace()()
}

func main() {
	A1()
	fmt.Println("ok")
}
