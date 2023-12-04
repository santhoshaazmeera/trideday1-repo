package main

import "fmt"

func main() {

	x := 12
	y := &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(&y)
	*y = 90 ///changing the value of x using its address .as x address is stored in y .
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(y) //here we only changed the value stored in y .as y is storing the addres of x .so the y we be same it wont change.that is the address of x
	fmt.Println(*y)
	fmt.Println(&y)
	z := 9
	y = &z
	fmt.Println(y)
	fmt.Println(*y)
}
