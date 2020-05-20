package main

import (
	"fmt"

	"github.com/nursabd/golanglearning/sample/nested/hello"
	"github.com/nursabd/golanglearning/sample/nested/say"
)

func main() {
	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}
