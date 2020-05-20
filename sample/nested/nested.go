package main

import (
	"Gotut/sample/nested/hello"
	"Gotut/sample/nested/say"
	"fmt"
)

func main() {
	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}
