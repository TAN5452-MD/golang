package main

import (
	"fmt"
	"strings"
)

func main() {
	Adder := MakeAddSuffix(".jpg")
	fmt.Printf("%v", Adder("罗俊吃盖饭"))
}

func MakeAddSuffix(pre string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, pre) {
			return name + pre
		}
		return name
	}
}
