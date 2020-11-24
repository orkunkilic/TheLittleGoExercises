package main

import "fmt"

type Ninja struct {
	Name string
	Age  int
}

func (s *Ninja) Super() {
	fmt.Println("Now I belong to Ninja", s.Name)
}

func main() {

	WooHoo := Ninja{"WooHoo", 32}
	fmt.Println(WooHoo)
	WooHoo.Super()
	//or
	WooHaa := Ninja{Name: "WooHaa", Age: 52}
	fmt.Println(WooHaa)
	WooHaa.Super()

	PointerHaa := &Ninja{"PointerHaa", 20}
	NonPointerHaa := Ninja{"NonPointerHaa", 20}
	increaseAgeP(PointerHaa)
	fmt.Println(PointerHaa)
	increaseAge(NonPointerHaa)
	fmt.Println(NonPointerHaa)

	EmptyHaa := new(Ninja) //or &Ninja{}
	fmt.Println(EmptyHaa)

}

func increaseAgeP(s *Ninja) {
	s.Age += 2
}

func increaseAge(s Ninja) {
	s.Age += 2
}
