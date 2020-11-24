package main

import "fmt"

type Person struct {
	Name string
}

//If my Ninja has an Introduce method, It will be default, this will be .Person.Introduce()
func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Ninja struct {
	*Person
	Age int
}

func main() {

	// and to use it:
	goku := &Ninja{
		Person: &Person{"Goku"},
		Age:    26,
	}
	goku.Introduce()
	goku.Person.Introduce()
}
