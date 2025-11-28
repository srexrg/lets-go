package main

import (
	"fmt"
)

func main (){

	nos := [6]int {1,2,3,4,5,6}

	for i:=0;i<len(nos);i++ {
		fmt.Println(nos[i])
	}
}