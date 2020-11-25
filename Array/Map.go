package main

import "fmt"

func main() {
	lookup := make(map[string]int, 5) //last argument is pre-definded len. Optional.
	lookup["go!"] = 9001
	power, exists := lookup["vegeta"]

	fmt.Println(power, exists)
	fmt.Println(len(lookup))
	delete(lookup, "go!")
	fmt.Println(lookup)
}
