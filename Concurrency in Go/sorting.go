package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sort"
)

func sortSlice(a []int, ch chan []int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("subarray to sort",a)

sort.Ints(a)
ch <- a
}
func main(){
	var wg sync.WaitGroup
	ch:= make( chan []int,4)
	var a []int
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 4 or more integers separated by space")
	scanner.Scan()
	input:=scanner.Text()
	parts := strings.Fields(input)
	for _, part := range parts{
		num, err :=strconv.Atoi(part)
		if err !=nil{
			fmt.Println("Error when converting", err)
			continue
		}
		a=append(a, num)
	}


	lengthOfTheSLice:=len(a)
	subArrayLength:= (lengthOfTheSLice/4)
	reminder:=lengthOfTheSLice%4
	
		// println("leght", lengthOfTheSLice)
		// println("subaraylenght", subArrayLength)
		// println("reminder", reminder)
wg.Add(4)
	firstArray:=a[0:subArrayLength]
	secondArray:= a[subArrayLength:subArrayLength*2]
	thirdArray:= a[(subArrayLength*2):subArrayLength*3]
	fourthArray:= a[(subArrayLength*3):(subArrayLength*4)+reminder]

    // fmt.Printf("fisrt array %v\n", firstArray)
	// fmt.Printf("second array %v\n", secondArray)
    // fmt.Printf("third array %v\n", thirdArray)
    // fmt.Printf("fourth array %v\n", fourthArray)


	go sortSlice(firstArray,ch,&wg)
	go sortSlice(secondArray,ch,&wg)
	go sortSlice(thirdArray,ch,&wg)
	go sortSlice(fourthArray,ch,&wg)
	go func() {
        wg.Wait()  
        close(ch)  
    }()

	var combined []int
	// fmt.Println(len(ch))
    for slice := range ch {
	combined=append(combined,slice...)
		sort.Ints(combined)
    }


					fmt.Printf("%v\n", combined)

}