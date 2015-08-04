package main

import (
	"fmt"

	. "github.com/GeorgeMac/funcy/functional"
)

func Output(list List) {
	fmt.Printf("%v", Head(list))
	if t := Tail(list); t != nil {
		fmt.Print(", ")
		Output(t)
	}
}

func add(x, y interface{}) interface{} {
	return x.(int) + y.(int)
}

func mul(i int, x interface{}) interface{} {
	return i * x.(int)
}

func double(x interface{}) interface{} { return mul(2, x) }

func square(x interface{}) interface{} { return mul(x.(int), x) }

func sum(list List) interface{} { return Reduce(add, list) }

func main() {
	fmt.Println("List")

	list := Cons(1, Cons(2, Cons(3, Cons(4, Cons(5, Cons(6, nil))))))
	Output(list)

	fmt.Println("\n\nDoubled List")

	list2 := Map(double, list)
	Output(list2)

	fmt.Println("\n\nMapped Square Function List")

	list3 := Map(square, list)
	Output(list3)

	fmt.Println("\n\nSummed (Reduced) Square Function List")

	list4 := Reduce(add, Map(square, list))
	fmt.Println(list4)
}
