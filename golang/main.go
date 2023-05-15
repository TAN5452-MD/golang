package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 1.23213
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))
	v := reflect.ValueOf(&x)
	v = v.Elem()
	v.SetFloat(2.3)
	fmt.Println(v)
}
