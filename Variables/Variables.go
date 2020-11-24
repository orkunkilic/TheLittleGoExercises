package main

import (
	"fmt"
)

func main() {
	myInt := 30
	myStr := "Hi, Dear"
	myFloat := 2.5

	fmt.Printf("This is my Int: %d\n", myInt)
	fmt.Printf("This is my Float: %.2f\n", myFloat)
	fmt.Printf("This is my String: %s\n", myStr)

	myInt = 55

	//myInt = "Can't do!"
	//myInt := 2000         "Can't also do!"

	fmt.Println("Hey, I do not need '%d'. Look->", myInt)
}
