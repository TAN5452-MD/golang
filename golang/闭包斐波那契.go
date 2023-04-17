package main

import "fmt"

func f() func(int) int {
	var x = 0
	var y = 0
	return func(j int) int {
		if j == 1 {
			x = 1
			return 1
		} else if j == 2 {
			y = 1
			return 1
		} else {
			temp := y
			y = x + y
			x = temp
			return temp
		}

	}
}
func main() {
	r := f()
	for i := 1; i < 10; i++ {
		fmt.Printf("%v\t", r(i))
	}
	//fmt.Println('s')
}
