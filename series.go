package main

import "fmt"

func main() {
	a := 1
	b := 1
	var c int
	fmt.Println("enter total no.of digit requied in series")
	var num int
	fmt.Scanln(&num)

	if num < 2 {
		fmt.Println("error")
	} else {

		fmt.Print(a)
		fmt.Print(",")
		fmt.Print(b)
		fmt.Print(",")

		for i := 3; i <= num; i++ {
			c = a + b
			fmt.Print(c)
			//fmt.Print(",")
			a = b
			b = c
		}
	}

}
