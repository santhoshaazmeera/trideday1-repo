package main

import "fmt"

func main() {
	print("ism main function")
	ad, s, m := maths(2, 3)
	fmt.Printf("the addition is : %d \n the subtraction is : %d \n the multiplication is %d ", ad, s, m)
	var a [5]int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}
	var summ = sum(a)
	fmt.Println(summ)

}
func maths(x, y int) (int, int, int) {
	var a int = x
	var b int = y
	var add int = a + b
	var sub int = a - b
	var mul int = a * b
	return add, sub, mul
}

func sum(a [5]int) int {
	var sum int = 0
	for _, value := range a {
		sum = sum + value
	}
	return sum

}
