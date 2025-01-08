package main

import(
	"fmt"
	"strconv"
)
func main(){
	var scanned string
	var slice []int
	for i:=0;i<10;i++{
			fmt.Println("Type an integer and press enter:")

	fmt.Scan(&scanned)
	integerInput, err:= strconv.Atoi(scanned)
	if err!=nil{
		fmt.Println("Error converting to integer")
	}
	slice=append(slice,integerInput)
	}
	BubbleSort(slice)
	fmt.Print(slice)
}

func Swap(slice []int, i int){
a:=slice[i]
slice[i]=slice[i+1]
slice[i+1]=a
}

func BubbleSort(slice []int){
	 var swapped bool= false
	n:=len(slice)
	for i:=0; i<n;i++{
		
		for j:=0; j< (n-i-1);j++{
			if slice[j]>slice[j+1]{
			Swap(slice,j)
			swapped=true
		}
		}
		if swapped!=true{
			return
		}

	}
}