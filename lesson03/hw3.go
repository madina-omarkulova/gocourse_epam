package main

import (
	"fmt"
	"sort"
)

func averv(a []int) float64 {
	s := 0.0
	for _, k := range a {
		s += float64(k)
	}
	return s / float64(len(a))
}

func max(s []string) string {
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
	var ans []int64
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
