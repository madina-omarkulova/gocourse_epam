package main

import (
	"fmt"

	"github.com/madina-omarkulova/gocourse_epam/fib"
)

func main() {
	var a int
	defer fmt.Println("Task completed.")
	fmt.Println("Please enter the number:")
	fmt.Scan(&a)
	fib.Printer(fib.Fibb(a))
}
