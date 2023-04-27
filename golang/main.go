package main

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
	var a = new(Simple)
	a.num = 1
	var b Simpler
	if _, ok := b.(a); ok {
		return
	}
}
