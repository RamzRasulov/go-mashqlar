package main

import "fmt"

func juftYokiToq(a int) string {
	if a%2 == 0 {
		return "juft"
	}
	return "toq"
}

func main() {
	var a int
	fmt.Scan(&a)

	fmt.Println(juftYokiToq(a))
}
