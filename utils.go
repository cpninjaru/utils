package utils

import "fmt"

func Contains(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	fmt.Println("Not found")
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	fmt.Println("Not found")
	return false
}
