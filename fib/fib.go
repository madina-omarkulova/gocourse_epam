package fib

import "fmt"

func Fibb(n int) {
	a, b, c := 0, 1, 1
	if n == 1 {
		fmt.Print("0 ")
	} else if n == 2 {
		fmt.Print("1  ")
	}
	for i := 2; i < n; i++ {
		fmt.Print(c)
		a, b, c = b, c, a+b
	}
}
