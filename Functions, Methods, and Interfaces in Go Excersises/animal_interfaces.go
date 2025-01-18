package main
import("fmt")

type Animal interface {
	Move()
	Eat()
	Speak()

}
type Snake struct{
	name string
	 food string
  locomotion string
  noise string
}

func (a Snake) Eat(){
  fmt.Println(a.food)
  }
func (a Snake) Speak(){
  fmt.Println(a.noise)
  }
func (a Snake) Move(){
  fmt.Println(a.locomotion)
  }



  type Cow struct{
	  name string
	 food string
  locomotion string
  noise string
}

func (a Cow) Eat(){
  fmt.Println(a.food)
  }
func (a Cow) Speak(){
  fmt.Println(a.noise)
  }
func (a Cow) Move(){
  fmt.Println(a.locomotion)
  }
  type Bird struct{
	  name string
	 food string
  locomotion string
  noise string
}

func (a Bird) Eat(){
  fmt.Println(a.food)
  }
func (a Bird) Speak(){
  fmt.Println(a.noise)
  }
func (a Bird) Move(){
  fmt.Println(a.locomotion)
  }

func NewBird(inputName string) Bird {
	return Bird{
		name:inputName,
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
}
  func NewSnake(inputName string) Snake {
	return Snake{
		name:inputName,
		food:       "mice",
		locomotion: "slither",
		noise:      "hss",
	}
}


  func NewCow(inputName string) Cow {
	return Cow{
		name:inputName,
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}
func main(){

var inputCommand string
var inputSecond string
var inputThird string
animals := make(map[string]interface{})

for i:=0;i<4;{
fmt.Println("Input:  newanimal name type or query name action")
fmt.Print(">")
fmt.Scan(&inputCommand,&inputSecond,&inputThird)


if inputCommand=="newanimal"{
	
	if inputThird=="bird" {
		bird:=NewBird(inputSecond)
		    animals[inputSecond] = bird
	} else if inputThird=="snake" {
		snake:=NewSnake(inputSecond)
			animals[inputSecond] = snake


	} else if inputThird=="cow" {
		cow:=NewCow(inputSecond)
		animals[inputSecond] = cow
	} else{
			fmt.Println("Type not found")
	}
	fmt.Println("Created it!")

} else if inputCommand=="query"{


    value, exists := animals[inputSecond]

    if exists {
		value, ok := value.(Animal)
		if ok{
	if inputThird=="speak"{
			value.Speak()
		} else if inputThird=="eat"{
			value.Eat()
		} else if inputThird=="move"{
			value.Move()
		}else{
           fmt.Println("Action not found!")
		}
		} else{
			           fmt.Println("Interface not satisfied")

		}
	
    } else {
        fmt.Printf("Animal '%s' does not exist", inputSecond)
    }


} else {
	print("Command not found")
}



println("")
}

  
}
