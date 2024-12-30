package main

import "fmt"

var m = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if _, ok := m[n]; ok {
		return m[n]
	}
	if n <= 2 {
		return 1
	}

	m[n] = fib((n - 1)) + fib((n - 2))
	return m[n]
}

func main() {

	fmt.Println(fib(90))

}
