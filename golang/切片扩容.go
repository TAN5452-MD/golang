package main

import "fmt"
/* 
给定 slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。
 */
func main() {
	var s = []int{1, 2, 3}
	magnify_slice(s, 5)
}

func magnify_slice(s []int, factor int) {
	fmt.Println(len(s))
	var newSlice = make([]int, len(s)*factor)
	copy(newSlice, s)
	fmt.Println(newSlice, len(newSlice))

}
