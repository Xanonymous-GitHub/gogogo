package main

import "fmt"

//幹你的XD
func fibonacci() func(int) int {
	return func(x int) int {
		a, b := -1, 1
		for i := 0; i <= x; i++ {
			tmp := b
			b += a
			a = tmp
		}
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
