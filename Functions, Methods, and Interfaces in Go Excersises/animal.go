package main
import("fmt")

type Animal struct {
  food string
  locomotion string
  noise string
}


func (a Animal) Eat(){
  fmt.Println(a.food)
  }
func (a Animal) Speak(){
  fmt.Println(a.noise)
  }
func (a Animal) Move(){
  fmt.Println(a.locomotion)
  }
func main(){
  cow:= Animal{
    food: "grass",
    locomotion: "walk",
    noise: "moo",
  }
    bird:= Animal{
    food: "worms",
    locomotion: "fly",
    noise: "peep",
  }
    snake:= Animal{
    food: "mice",
    locomotion: "slither",
    noise: "hsss",
  }
var inputName string
var inputAction string
for i:=0;i<4;{
fmt.Println("Input an animal name and one action separated by a space")
fmt.Print(">")
fmt.Scan(&inputName,&inputAction)
var a Animal
if inputName=="cow"{
	a=cow
} else if inputName=="bird"{
	a=bird
} else if inputName=="snake"{
	a=snake
} else {
	print("Animal not found")
}


if inputAction=="eat"{
	a.Eat()
} else if inputAction=="move"{
	a.Move()
} else if inputAction=="speak"{
	a.Speak()
}else{
	print("Action not found")
}
println("")
}

  
}


