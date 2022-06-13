package main

import "fmt"

func main() {
	a := 1
	b := 1
	var c int
	fmt.Println("enter total no.of digit requied in series")
	var num int
	fmt.Scanln(&num)

	fmt.Print(a, ", ", b, ", ")
	for i := 2; i < num; i++ {
		c = a + b
		a = b
		b = c
		if i == num-1 {
			fmt.Print(c, "\n")
		} else {
			fmt.Print(c, ", ")
		}
	}

}
