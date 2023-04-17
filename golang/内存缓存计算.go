package main

import (
	"fmt"
	"time"
)

const LIM = 41

// 该语句的作用是提前在内存中声明内存区域
// 声明一个41位的unit64的数组
var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	//获取2个时间段的差值
	delta := end.Sub(start)
	fmt.Printf("fibonacci took this amount of time: %s\n", delta)

}
func fibonacci(n int) (res uint64) {
	// 检查当前位是否计算过
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
