# Go Mirror
![Golang Mirror](image/SpidermanPoint.jpg)

## Preface
Well I need some way to copy data between struct without having to assign it one by one. One solution that i found is by encoding it to json then decode it in the target struct. But I think it's a hacky solution. Also, I need to convert arbitrary data to other data type but I seems can't find one. I already tried searching the internet for non hacky solution but seems unable to find one. So I decided to create one by myself. 

## What this thing do
Convert different struct to another struct or map or vice versa. It utilize reflection to achieve it's target

## Usage
```golang
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
	fmt.Println("Raw Map : ", rawMap)
	//OUTPUT : Raw Map :  map[Age:2 Child:map[Age:1 Name:X-DORU Species:404] Name:Doruru Species:Digital Or Unknown]
	//Lets see our final result
	fmt.Println("New Organism : ", parentOrganism)
	//OUTPUT : New Organism :  {Doruru 2 Digital Or Unknown {X-DORU 1 404}}
}
```

## Benchmark
Let's see the performance result
```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/go-mirror
BenchmarkJson-8          	  279831	      4048 ns/op	     496 B/op	      12 allocs/op
BenchmarkJsoniter-8      	  601572	      1985 ns/op	     296 B/op	      14 allocs/op
BenchmarkMirror-8        	  923041	      1292 ns/op	     192 B/op	      15 allocs/op
BenchmarkSmartMirror-8   	  925382	      1310 ns/op	     192 B/op	      15 allocs/op
PASS
ok  	github.com/firmanmm/go-mirror	5.266s
```
Eventhough this package do perform faster the `hacky` methods, it doesn't have much support than them. One of the downside of this library is that it currently doesn't support using pointer.
## Todo
- Create more Example
- Add support to pointer
- Convert switch to jump table
- More Test
- Benchmark