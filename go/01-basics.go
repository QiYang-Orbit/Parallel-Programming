package main

import (
	"fmt"
)

func main() {
	fmt.Println("hii")
	fmt.Println(7/3)

	var a = "hello"
	var b, c int = 1, 2
	fmt.Println(a, b, c)

	var d = true
	fmt.Println(d)

	e := "hello"
	fmt.Println(e)

	e = "modified e"
	fmt.Println(e)

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	fmt.Println(s)

	s = append(s, "e", "f")
	fmt.Println(s)
	

}