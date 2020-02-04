package main

import "fmt"

func fibonacci() func(int) int {
	a, b := -1, 1
	return func(x int) int {
		for i := 0; i <= i; i++ {
			tmp := b
			b += a
			a = tmp
		}
		fmt.Println(b)
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
		fmt.Println(i)
	}
}
