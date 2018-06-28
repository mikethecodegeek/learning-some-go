package main

import "fmt"

func fibb(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	}

	return fibb(n-1) + fibb(n-2)

}

func main () {
	fmt.Println(fibb(10))
}