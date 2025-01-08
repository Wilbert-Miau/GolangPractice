package main

import (
	"strings"
	"fmt"
	"regexp"
	"bufio"
	"os"
)

func main(){
	pattern := `^[iI][a-zA-Z\s]*a+[a-zA-Z\s]*[nN]$`
	re := regexp.MustCompile(pattern)
	i:=0
	for i<5 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string: ")
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)
	if re.MatchString(inputString){
	fmt.Println("Found!")

	} else
	{
		fmt.Println("Not Found!")

	}

	}

}