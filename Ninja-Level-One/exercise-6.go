package main

import (
	"fmt"
)
const currentYear = 2020
const (
	a = currentYear + iota
	b = currentYear + iota
	c = currentYear + iota
	d = currentYear + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
