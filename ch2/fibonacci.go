/*
 * @program: The-Go-Programming-Language
 * @author: Leon
 * @create: 2020-09-13 22:33
 */
package main

func fin(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
