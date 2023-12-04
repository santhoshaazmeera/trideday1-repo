package main

import "fmt"

func main() {
	var s string
	fmt.Println("enter string :")
	fmt.Scan(&s)
	var chareccount [26]int
	//var newarr []byte
	for j := 0; j < len(s); j++ {
		var i rune = rune(s[j]) - rune('a')
		fmt.Println(i)
		chareccount[i] = chareccount[i] + 1
		fmt.Println(chareccount[i])
	}
	// var u int = 0
	// for j := 0; j < len(s); j++ {

	// 	for _, k := range newarr {
	// 		fmt.Println(s[j] == k)
	// 		if s[j] == k {
	// 			break

	// 		} else {

	// 			newarr[u] = s[j]
	// 			u = u + 1

	// 			//fmt.Println(u, newarr[u])
	// 			break
	// 		}
	// 	}

	// }

	// for i := 0; i < len(newarr); i++ {
	// 	fmt.Printf("%c count : %d \n", newarr[i], chareccount[rune(newarr[i])-rune('a')])
	// }

}
