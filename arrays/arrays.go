// Package main provides ...
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter a list of comma separated numbers: ")

	var input string
	fmt.Scanf("%s", &input)
	var numbers []int

	for _, val := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}

	result := sum(numbers) // 15
	fmt.Println(result)
}

func sum(array []int) int {
	total := 0
	for _, num := range array {
		total += num
	}
	return total
}
