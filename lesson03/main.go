package main

import (
	"fmt"
	"sort"
)

func main() {
	defer fmt.Println("\nDone. \n ")
	a := []int{1, 2, 3, 4, 5, 6}
	//2 examples for an array function:
	fmt.Println("Input for array function:", a)
	a2 := averv(a)
	fmt.Println("Output: ", a2)
	b := []int{1, -2, 3, -5, 9, 10, -16}
	fmt.Println("Input for array function:", b)
	b2 := averv(b)
	fmt.Println("Output: ", b2)
	//Now slices max function:
	s := []string{"one", "two", "three"}
	fmt.Println("\nInput for max function: ", s)
	s2 := max(s)
	fmt.Println("Output: ", s2)
	l := []string{"z", "a", "e", "a"}
	fmt.Println("Input for max function: ", l)
	l2 := max(l)
	fmt.Println("Output: ", l2)
	fmt.Print("\n")

	//Slices reverse function:
	k := []int64{1, 2, 3, 4, 5, 6}
	fmt.Println("\nInput for reverse function: ", k)
	k2 := reverse(k)
	fmt.Println("Output: ", k2)
	n := []int64{1, -2, 3, -5, 9, 10, -16}
	fmt.Println("Input for reverse function: ", n)
	n2 := reverse(n)
	fmt.Println("Output: ", n2)
	//map printSorted function:
	m := map[int]string{2: "a", 0: "b", 1: "c"}
	fmt.Print("\nInput for printSorted function: ", m, "\nOutput: ")
	printSorted(m)
	p := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Print("Input for printSorted function: ", p, "\nOutput: ")
	printSorted(p)
}

func averv(a []int) float64 {
	if len(a) == 0 {
		return 0
	}
	s := 0.0
	for _, k := range a {
		s += float64(k)
	}
	return s / float64(len(a))
}

func max(s []string) string {
	if len(s) == 0 {
		return ""
	}
	ans := s[0]
	mx := len(s[0])
	for _, k := range s {
		if len(k) > mx {
			ans = k
			mx = len(k)
		}
	}
	return ans
}

func reverse(k []int64) []int64 {
	ans := make([]int64, 0, len(k))
	for i := len(k) - 1; i > -1; i-- {
		ans = append(ans, k[i])
	}
	return ans
}

func printSorted(m map[int]string) {
	n := make([]int, 0, len(m))
	for k := range m {
		n = append(n, k)
	}
	fmt.Print("[")
	sort.Ints(n)
	for i, l := range n {
		fmt.Print("\"", m[l], "\"")
		if i != len(n)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
