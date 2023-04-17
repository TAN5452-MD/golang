package main

import "fmt"

/*
使用高阶函数对一个集合进行过滤：s 是前 10 个整数的一个切片。建立一个函数 Filter，
它接受 s 作为第一参数，fn func (int) bool 作为第二参数，并返回满足函数 fn 的 s 元素的切片（使其为真）。
用 fn 测试整数是否为偶数。
*/
func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	r := Filter(s, Judge)
	fmt.Println(r())

}
func Filter(s []int, fn func(int) bool) func() []int {
	var slice = make([]int, len(s))
	return func() []int {
		for _, v := range s {
			if fn(v) {
				slice = append(slice, v)
			}
		}
		return slice
	}
}

func Judge(number int) bool {
	return number%2 == 0
}
