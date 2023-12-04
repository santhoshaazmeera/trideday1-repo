package main
import "fmt"
import "unicode/utf8"
func main(){
	fmt.Println("hello golang bye no no hi!!")
     fmt.Println("D")
    var s string = `hello23
worlo
k`
fmt.Println(len(s))
fmt.Println(utf8.RuneCountInString(s))
//var t string ="క ఖ గ ఘ ఙ"
//    fmt.Println("telugu")
//    
//    fmt.Println(len(t))

    var rr int = utf8.RuneCountInString(s)
    fmt.Println(rr)
    var r rune ='a'
    fmt.Println(r)
}
