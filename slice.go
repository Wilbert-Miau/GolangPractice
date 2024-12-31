package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main(){
	slice := make([]int,3)
	i:=0
	for i<5 {
	fmt.Printf("Enter an integer: ")
	var inputString string

	fmt.Scan(&inputString)
	if inputString=="X" || inputString=="x"{
		return
	}
	in,err := strconv.Atoi(inputString)
	if err!=nil{
			fmt.Println("Error converting to integer")
			continue
	}
	slice=append(slice,in)
	sort.Ints(slice)
	fmt.Printf("Slice: %v\n",slice)

	}
}