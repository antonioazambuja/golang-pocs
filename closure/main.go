package main

import "fmt"

func testClosure() func() {
	anyNumber := 10
	return func() {
		fmt.Println(anyNumber)
	}
}

func main() {
	anyNumber := 20
	fmt.Println(anyNumber)
	anyNumberClosure := testClosure()
	anyNumberClosure()
}
