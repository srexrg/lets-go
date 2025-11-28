package main

import "fmt"

type Volt struct {
	X int
	Y int
}

func main (){

	a:=10
	b:=&a
	fmt.Println(a,b)
	fmt.Println(*b)

	*b=20
	fmt.Println(a,b)
	fmt.Println(*b)

	v :=Volt{8,10}

	u:= &v

	u.X = 20
	fmt.Println(v)


}