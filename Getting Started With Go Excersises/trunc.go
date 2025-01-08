package main

import (
	"fmt"
	"strconv"
)

func main(){
	i:=0
	for i<5 {
	var inputString string
	fmt.Printf("Enter one float number: ")
	fmt.Scan(&inputString)
	inputFloat, err := strconv.ParseFloat(inputString,32)
		if err!=nil{
		fmt.Println("conversion failed")
		return
	}
	inputInteger := int(inputFloat)

	outputString := strconv.Itoa(inputInteger)
	fmt.Println(outputString)
	}

}