package main

import "os"

func main() {
	f, err := os.Open("./b")
	if err != nil {
		os.Exit(1)
		return
	}
}
