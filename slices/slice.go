package main

import "fmt"


func main(){

	// s:= []int{1,2,3,4,5}

	// s=s[:0]
	// PrintSlice(s)
	
	// s=s[2:4]
	// PrintSlice(s)
	
	a := make([]int,5)
	PrintSlice(a)
	
	b:= make([]int,0,5)
	PrintSlice(b)
	
	c:= b[:2]
	PrintSlice(c)
	
	d:=c[2:5]
	PrintSlice(d)


}

func PrintSlice(s [] int){

	fmt.Printf("len %d cap %d slice %v \n",len(s),cap(s),s)
}