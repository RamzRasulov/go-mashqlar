package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a >= b && a >= c {
		fmt.Printf("katta son a=%d\n", a)
	} else if b >= a && b >= c {
		fmt.Printf("katta son b=%d\n", b)
	} else {
		fmt.Printf("katta son c=%d\n", c)
	}
}
