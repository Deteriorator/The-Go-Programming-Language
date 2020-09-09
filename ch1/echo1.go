package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	/*
	 * exam 1.3
	 */
	a := time.Now()
	fmt.Println(a)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	b := time.Since(a)
	fmt.Println(b)
	fmt.Println(s)
}
