package main

import "fmt"

func main() {
	var base, height float64

	fmt.Scanln(&base)
	fmt.Scanln(&height)

	fmt.Println((base * height) / 2)
}
