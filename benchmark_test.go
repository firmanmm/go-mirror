package mirror

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type Organism struct {
	Name    string
	Age     uint
	Species string
}

type ParentOrganism struct {
	Name    string
	Age     uint
	Species string
	Child   Organism
}

func BenchmarkJson(b *testing.B) {
	source := ParentOrganism{
		Name:    "Rendoru",
		Age:     22,
		Species: "Human",
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		body, err := json.Marshal(&source)
		if err != nil {
			b.Error(err.Error())
		}
		dest := ParentOrganism{}
		err = json.Unmarshal(body, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkJsoniter(b *testing.B) {
	source := ParentOrganism{
		Name:    "Rendoru",
		Age:     22,
		Species: "Human",
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		body, err := jsoniter.Marshal(&source)
		if err != nil {
			b.Error(err.Error())
		}
		dest := ParentOrganism{}
		err = jsoniter.Unmarshal(body, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkMirror(b *testing.B) {
	source := ParentOrganism{
		Name:    "Rendoru",
		Age:     22,
		Species: "Human",
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := ParentOrganism{}
		err := Mirror(&source, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkSmartMirror(b *testing.B) {
	source := ParentOrganism{
		Name:    "Rendoru",
		Age:     22,
		Species: "Human",
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := ParentOrganism{}
		err := SmartMirror(&source, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}
