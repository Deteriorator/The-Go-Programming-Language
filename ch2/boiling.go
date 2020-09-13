/*
 * @program: The-Go-Programming-Language
 * @author: Leon
 * @create: 2020-09-13 22:02
 */
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}
