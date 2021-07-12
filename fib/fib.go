package fib

import "fmt"

func Printer(n int) {
	a, b, c := 0, 1, 1
	fmt.Print("0 1 ")
	for c <= n {
		fmt.Print(c, " ")
		a = b
		b = c
		c = a + b
	}
	fmt.Print("\n")
}

func Fibb(n int) int {
	b, c := 1, 1
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	for i := 2; i < n; i++ {
		b, c = c, c+b
	}
	return c
}
