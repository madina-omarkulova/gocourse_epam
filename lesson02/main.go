package main

import (
	"fmt"

	"github.com/madina-omarkulova/gocourse_epam/lesson02/les02"
)

func main() {
	var a int
	defer fmt.Println("Task completed.")
	fmt.Println("Please enter the number:")
	fmt.Scan(&a)
	les02.Fibb(a)
}
