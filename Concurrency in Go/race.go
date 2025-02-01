package main

import "fmt"

func add1000(a *int){
	for i:=0;  i< 100000;i++{
		*a=*a+1
	}
}
func main(){
	a:=0

	go add1000(&a)
	go add1000(&a)
	go add1000(&a)
	fmt.Println("this will 300000?")

	fmt.Println(a)
}