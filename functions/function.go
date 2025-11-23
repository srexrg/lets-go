package main

import (
	"fmt"
)

var var1 ,var2 int

var ao = 0


func add(x int,y int) int {
	return x+y
}

func swap(a,b string) (string,string){

	return b,a
}


func main(){
	newvar := 10
	fmt.Println(add(1,2))
	a,b := swap("world","hello")
	fmt.Println(a,b)
	fmt.Println(newvar,ao,var1,var2)

}