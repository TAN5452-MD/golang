package main

import (
	"fmt"
	"strings"
)

func main() {
	//判断字符串以什么开头
	var s string = "qwesaffe"
	fmt.Printf("%t", strings.HasPrefix(s, "fa"))

	//判断字符串以什么结尾
	fmt.Printf("%t",strings.HasSuffix(s,"ffe"))
}
