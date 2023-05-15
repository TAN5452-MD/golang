package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(int)
}
type Simple struct {
	num int
}

func (s Simple) Get() int {
	return s.num
}

// 如果不使用引用的话只是值传递并不能真正修改到值
func (s *Simple) Set(num int) {
	s.num = num
}

func main() {
	var a Simpler
	sq := new(Simple)
	a = sq
	if v, ok := a.(*Simple); ok {
		fmt.Println(v)
		fmt.Printf("%v", ok)
	}
}
