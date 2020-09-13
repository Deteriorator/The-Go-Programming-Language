/*
 * @program: The-Go-Programming-Language
 * @author: Leon
 * @create: 2020-09-13 22:21
 */
package main

import (
	"fmt"
	"os"
	"strconv"
)

func GCD(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// there is problem
func main() {
	a := os.Args[1]
	b := os.Args[2]
	fmt.Printf(string(GCD(strconv.Atoi(a), strconv.Atoi(b))))
}
