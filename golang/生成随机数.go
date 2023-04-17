package main

import (
	"math/rand"
)

func main() {
	random()
}
func random() {
	for i := 0; i < 10; i++ {
		a := rand.Intn(8)
		println(a)
		println(a)
	}
}
