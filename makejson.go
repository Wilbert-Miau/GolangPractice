package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
		"encoding/json"

)

func main() {
infoMap:=make(map[string]string)
fmt.Println("Enter a name")
reader := bufio.NewReader(os.Stdin)
inputString, _ := reader.ReadString('\n')
inputString=strings.TrimSpace(inputString)
infoMap["name"]=inputString


fmt.Println("Enter the address")
inputString, _ = reader.ReadString('\n')

inputString=strings.TrimSpace(inputString)
infoMap["address"]=inputString

jsonData,err:= json.Marshal(infoMap)
if err != nil{
			fmt.Println("Error converting to JSON:", err)
			return
}

fmt.Println("json: ", string(jsonData))
}