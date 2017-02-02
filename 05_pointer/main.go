package main

import "fmt"

func main() {
	a := 7
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	*b++
	fmt.Println(a)
}
