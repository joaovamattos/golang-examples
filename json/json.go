package main

import (
	"fmt"
    "encoding/json"
)

type Pokemon struct {
    Name string
    Moves []string
}

func main() {
	// Struct to JSON
	charmander := &Pokemon{
		Name: "Charmander",
		Moves:[]string{"Scratch", "Growl", "Ember", "Smokescreen"}}    
	fmt.Println("struct", charmander)
	jsonCharmander, _ := json.Marshal(charmander)	
	fmt.Println("json", string(jsonCharmander))

	// JSON to Struct
	chameleon := `{"Name":"Charmeleon","Moves":["Scratch", "Ember", "Metal Claw", "Flamethrower"]}`
    fmt.Println("json", chameleon)
    structCharmeleon := Pokemon{}
    json.Unmarshal([]byte(chameleon), &structCharmeleon)
    fmt.Println("struct", structCharmeleon)
}