package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0]) // exam 1.1
	fmt.Println(len(os.Args))
	a := time.Now()
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, " ", os.Args[i]) //exam 1.2
	}
	b := time.Since(a)
	fmt.Println(b)
}
