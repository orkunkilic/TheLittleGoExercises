package main

import "fmt"

func main() {

	a := 3
	myPower := power(a)
	fmt.Println("Power of", a, "is", myPower)

}

func power(a int) int {
	return a * a
}
