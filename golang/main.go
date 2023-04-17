package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type T struct {
	a int
	b float32
	c string
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Println(t)
}

func (t *T) String() string {
	var pat = "\t"
	re := regexp.MustCompile(pat)
	str := re.ReplaceAllString(t.c, "\\t")
	return "" + strconv.Itoa(t.a) + "/" + strconv.FormatFloat(float64(t.b), 'g', -1, 32) + "/" + str + ""
}
