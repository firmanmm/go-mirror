package main

import (
	"fmt"
	"log"

	"github.com/firmanmm/go-mirror"
)

type Person struct {
	Name string
	Age  int
}

type Organism struct {
	Name    string
	Age     int
	Species string
}

func main() {
	//Lets initialize person as our source
	person := Person{
		Name: "Rendoru",
		Age:  22,
	}
	//Now lest create empty organism
	organism := Organism{}
	//Here we want to copy (Mirror) Person's data to Organism
	if err := mirror.Mirror(&person, &organism); err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Person : ", person)
	//OUTPUT : Person :  {Rendoru 22}
	//Since person didn't have Species, organism species will remain empty
	fmt.Println("Organism :", organism)
	//OUTPUT : Organism : {Rendoru 22 }
	//Lets set species for our organism
	organism.Species = "Human"
	//Now let's try to copy (Mirror) organism to map
	organismData := map[string]interface{}{}

	if err := mirror.Mirror(&organism, &organismData); err != nil {
		log.Fatalln(err.Error())
	}
	//Lets print our organism first
	fmt.Println("Organism : ", organism)
	//OUTPUT : {Rendoru 22 Human}
	//Lets see our organism map data
	fmt.Println("Organism Map : ", organismData)
	//OUTPUT : Organism Map :  map[Age:22 Name:Rendoru Species:Human]
	//Lets convert back our organism data to struct
	newOrganism := Organism{}
	if err := mirror.Mirror(&organismData, &newOrganism); err != nil {
		log.Fatalln(err.Error())
	}
	//Lets print our new organism
	fmt.Println("New Organism : ", newOrganism)
	//OUTPUT : New Organism :  {Rendoru 22 Human}

	//How About We Take it to the next level?
	//Lets use Incorrect data type
	rawMap := map[string]interface{}{
		"Name":    "Doruru",
		"Age":     "2", //This is a string and not a number
		"Species": "Digital Or Unknown",
		"Child": map[string]interface{}{
			"Name":    "X-DORU",
			"Age":     float64(1.000), //This is a float and not an int
			"Species": 404,            //This is a number and not a string
		},
	}

	//Lets define a new local type
	type ParentOrganism struct {
		Name    string
		Age     uint //Lets try Uint
		Species string
		Child   Organism
	}

	parentOrganism := ParentOrganism{}
	//Since we are dealing with different data type, using Mirror will cause error
	if err := mirror.Mirror(&rawMap, &parentOrganism); err != nil {
		fmt.Println(err.Error()) //lets temporary change it to println
		//OUTPUT : Destination field type didn't match Source field type
	}
	//Now let's use SmartMirror, It behaves like mirror but perform transformation as needed
	//Also, it won't stop on error
	if err := mirror.SmartMirror(&rawMap, &parentOrganism); err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Raw Map : ", rawMap)
	//OUTPUT : Raw Map :  map[Age:2 Child:map[Age:1 Name:X-DORU Species:404] Name:Doruru Species:Digital Or Unknown]
	//Lets see our final result
	log.Println("New Organism : ", parentOrganism)
	//OUTPUT : {Doruru 2 Digital Or Unknown {X-DORU 1 404}}
}
