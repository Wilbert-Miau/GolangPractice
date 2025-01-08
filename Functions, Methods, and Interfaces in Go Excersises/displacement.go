package main

import (
"fmt"
"strconv"
"math"
)

func main (){
var inputAceleration string
var inputInitialDisplacement string
var inputVelocity string

fmt.Println("Enter aceleration")
fmt.Scan(&inputAceleration)
aceleration,err:= strconv.ParseFloat(inputAceleration,64)
fmt.Println("Enter initial displacement")
fmt.Scan(&inputInitialDisplacement)
fmt.Println("Enter velocity")
fmt.Scan(&inputVelocity)



initialDisplacement, err:= strconv.ParseFloat(inputInitialDisplacement,64)
velocity,err:=strconv.ParseFloat(inputVelocity,64)

if err!=nil{
	fmt.Println("Error when trying to convert", err)
}
fn:=GenDisplaceFn(aceleration,velocity,initialDisplacement)


var inputTime string
fmt.Println("Enter time")
fmt.Scan(&inputTime)
time, err:= strconv.ParseFloat(inputTime,64)

if err!=nil{
	fmt.Println("Error when trying to convert", err)
}

fmt.Println("displacement after input time:",fn(time))
fmt.Println("displacement after 5s:",fn(5))
fmt.Println("displacement after 3s:",fn(3))



}


func GenDisplaceFn( a float64,  vo float64,  so float64) func(float64) float64 {


	return func (t float64) float64{
		s1:=(a*(math.Pow(t,2)))/2 + vo*t + so
		return s1
	}
}