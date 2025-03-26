package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func modifyVal(slice []string) {
	slice[0] = "modified"
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	mySlice := []string{"hello", "world"}
	fmt.Println("Before modification:", mySlice)
	modifyVal(mySlice)
	fmt.Println("After modification:", mySlice)

}