// Package main provides
package main

import "fmt"

func main() {
	x, y := 200, 4
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x, y *int) {
	temp := *x

	*x = *y
	*y = temp
}
