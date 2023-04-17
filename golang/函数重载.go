package main

import (
	"fmt"
)

func main()  {
   type fn func(a int, b int) int
   var function fn 	
   function = add
	r := function(1,2)
	fmt.Println(r)
	function = max
	r2 := function(1,2)
	fmt.Println(r2)
}

func add(a int , b int ) int {
	return a+b
}

func max(a int, b int) int {
	return int(a>b)
}