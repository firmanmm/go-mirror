# Go Mirror
![Golang Mirror](image/SpidermanPoint.jpg)

## Preface
Well I need some way to copy data between struct without having to assign it one by one. One solution that i found is by encoding it to json then decode it in the target struct. But I think it's a hacky solution. Also, I need to convert arbitrary data to other data type but I seems can't find one. I already tried searching the internet for non hacky solution but seems unable to find one. So I decided to create one by myself. 

## About
Convert different struct to another struct or map or vice versa. It utilize reflection to achieve it's target. This library will try to not perform duplication if possible, so some data such as pointer will be copied by their pointer value and not the pointed data.
Useful if you need a quick and easy way to convert one struc to other.

## Usage
Please see test and benchmark file for more example.
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
BenchmarkJsonStructToSameType-8                      	  133640	      9151 ns/op	     992 B/op	      18 allocs/op
BenchmarkJsoniterStructToSameType-8                  	  256162	      4827 ns/op	    1048 B/op	      29 allocs/op
BenchmarkMirrorStructToSameType-8                    	 8142504	       152 ns/op	     128 B/op	       1 allocs/op
BenchmarkSmartMirrorStructToSameType-8               	 8034150	       154 ns/op	     128 B/op	       1 allocs/op
BenchmarkJsonStructToOtherType-8                     	  139759	      8862 ns/op	     992 B/op	      18 allocs/op
BenchmarkJsoniterStructToOtherType-8                 	  267295	      4838 ns/op	    1048 B/op	      29 allocs/op
BenchmarkMirrorStructToOtherType-8                   	  584518	      2207 ns/op	     320 B/op	      20 allocs/op
BenchmarkSmartMirrorStructToOtherType-8              	  533973	      2208 ns/op	     320 B/op	      20 allocs/op
BenchmarkJsonStructToMap-8                           	   81439	     14154 ns/op	    3210 B/op	      75 allocs/op
BenchmarkJsoniterStructToMap-8                       	  149857	      8493 ns/op	    2819 B/op	      68 allocs/op
BenchmarkMirrorStructToMap-8                         	  201561	      5995 ns/op	    1617 B/op	      49 allocs/op
BenchmarkSmartMirrorStructToMap-8                    	  195078	      6079 ns/op	    1617 B/op	      49 allocs/op
BenchmarkJsonMapToStruct-8                           	   69858	     21069 ns/op	    3073 B/op	      65 allocs/op
BenchmarkJsoniterMapToStruct-8                       	  161488	      7295 ns/op	    1348 B/op	      36 allocs/op
BenchmarkMirrorMapToStruct-8                         	   56398	     19227 ns/op	    2832 B/op	     170 allocs/op
BenchmarkSmartMirrorMapToStruct-8                    	   60858	     18411 ns/op	    2832 B/op	     170 allocs/op
BenchmarkJsonStructToMapThenToOtherStruct-8          	   35740	     34084 ns/op	    6284 B/op	     140 allocs/op
BenchmarkJsoniterStructToMapThenToOtherStruct-8      	   75412	     15802 ns/op	    4174 B/op	     104 allocs/op
BenchmarkMirrorStructToMapThenToOtherStruct-8        	   44391	     24779 ns/op	    4449 B/op	     219 allocs/op
BenchmarkSmartMirrorStructToMapThenToOtherStruct-8   	   49772	     25747 ns/op	    4449 B/op	     219 allocs/op
PASS
ok  	github.com/firmanmm/go-mirror	29.776s
```
Eventhough this package do perform faster than the `hacky` methods on certain scenario, it doesn't duplicate the data by default. So if you are looking for duplication, please find other package.
## Todo
- Create more Example
- Improve conversion from map to struct
- More Test
- Benchmark