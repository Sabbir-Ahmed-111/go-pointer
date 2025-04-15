package main

import "fmt"

// pass by value
// pass by reference

func point(numbers [3]int) {
	fmt.Println(numbers)
} //pass by value normal vabe kaj hoy,point a na

func add(num *[3]string) {
	fmt.Println(num) //pass by reference aita pointer ar kaj
}

func main() {
	//pointer or address memory (ram)
	/*
		x := 20

		p := &x //& => ampersand - address of

		fmt.Println(x)
		fmt.Println(p) //print value ar addr	of
		fmt.Println(*p) //Value of address
	*/

	arr := [3]int{1, 2, 3}
	point(arr) //pass by value

	arr2 := [3]string{"I", "Love", "You"}
	add(&arr2) //pass by reference aita pointer ar kaj

}
