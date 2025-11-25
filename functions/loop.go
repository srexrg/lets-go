package main

import(
	"fmt"
)

func Loop (a,b int) int{

	for i:=a;i<=b;i++{
		a+=b
	}

	fmt.Println("a",a)

	return a


}