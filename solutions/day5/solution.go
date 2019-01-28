package main

import "fmt"

type Pair struct {
	a int
	b int
}

func main() {
	fmt.Println(car(cons(3, 4)))
	fmt.Println(cdr(cons(3, 4)))
}

func cons(a int, b int) Pair {
	return Pair{a, b}
}

func car(pair Pair) int {
	return pair.a
}

func cdr(pair Pair) int {
	return pair.b
}
