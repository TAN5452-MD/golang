package main

import "fmt"

func main() {
	fmt.Print("1")
	fmt.Print("1")
	fmt.Print("1")
	fmt.Print("1")
}
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8000", nil)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style='color:orange'>我儿胡发秋速速学习</h1>")
}
