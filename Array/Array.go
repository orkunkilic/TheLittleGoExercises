package main

import "fmt"

func main() {
	var scores [10]int
	scores[0] = 339
	fmt.Println(scores)

	for index, value := range scores {
		fmt.Println("Index:", index, "& Value:", value)
	}

	scoresSlice := make([]int, 0, 10)

	//scoresSlice[7] = 933 Error
	fmt.Println(scoresSlice)

	scoresSlice = append(scoresSlice, 5)
	fmt.Println(scoresSlice)

	scoresSlice = scoresSlice[0:8]
	scoresSlice[7] = 933
	fmt.Println(scoresSlice)
}
