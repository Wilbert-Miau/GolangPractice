package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
	FName string
	LName string
}
func main(){
	fmt.Println("giveme a textfile name")
	var filename string
	fmt.Scanln(&filename)
	file,err:=os.Open(filename)
	if err!=nil{
		fmt.Println("error at opening file")
		return
	}
	defer file.Close()
	var data []Name
	scanner:= bufio.NewScanner(file)
	for scanner.Scan(){
		line :=scanner.Text()
		words := strings.Fields(line)
		if len(words)!=2{
			fmt.Println("error")
		}
		nameData:= Name{
			FName :words[0],
			LName: words[1],
		}
		data = append(data,nameData)
	}

	for _,val:= range data{
		fmt.Printf("First Name: %s, Last Name: %s \n", val.FName,val.LName)
	}
}