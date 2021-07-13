package main

import (
	"fmt"

	"github.com/madina-omarkulova/gocourse_epam/lesson02/fib"
)

func main() {
	defer fmt.Println("Task is completed.")
	var a int
	fmt.Println("Please enter the number:")
	fmt.Scan(&a)
	fib.Printer(a)
}
