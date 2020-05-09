package main

import (
	"fmt"
)

type MyString string

type MyFunc func(string) int

func main() {
	testMyString()
	testMyFunc()
}

func testMyString() {
	a := MyString("testVal")
	fmt.Println("testMyString: ", a)
	addDivider()
}

func testMyFunc() {
	a := MyFunc(func(s string) int {
		fmt.Println("testFuncType: string passed = ", s)
		return 123
	})
	fmt.Printf("testFuncType: ret val = %v \n", a("zz"))
	addDivider()
}

func addDivider() {
	fmt.Println("-------------------------------------------")
}
