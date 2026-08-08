package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Son kiriting: ")
	fmt.Scan(&a)

	b = a % 3

	fmt.Println(b)
}
