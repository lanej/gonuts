package main

//import "fmt"

//func main() {
//x := 1.5
//square(&x)

//fmt.Println(x)
//}

func square(x *float64) {
	*x = *x * *x
}
