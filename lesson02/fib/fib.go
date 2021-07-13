package fib

import "fmt"

func Printer(n int) {
	for i := 0; i <= n; i++ {
		fmt.Print(Fibb(i), " ")
	}
	fmt.Print("\n")
}

func Fibb(n int) int {
	if n < 0 {
		return -1
	}
	if n < 2 {
		return n
	}
	return Fibb(n-1) + Fibb(n-2)
}
