package main

import (
	"fmt"
	"regexp"
)

func main() {
	const a = "abcdefg"
	pat := "abc"

	re, _ := regexp.Compile(pat)
	fmt.Println(re)

	//将匹配到的部分进行替换
	str := re.ReplaceAllString(a, "AAA")
	fmt.Println(str)

	//参数为函数时
	f := func(s string) string {
		return "AAA"
	}
	str2 := re.ReplaceAllStringFunc(a, f)
	fmt.Println(str2)
}
