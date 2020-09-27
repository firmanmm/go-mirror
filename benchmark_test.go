package mirror

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type Organism struct {
	Name    string
	Age     uint
	Species string
}

type ParentOrganism struct {
	Name         string
	Age          uint
	Species      string
	Active       bool
	Passive      bool
	Weight       float64
	Fingerprint  []byte
	Child        Organism
	PointerChild *Organism
}

type AlienOrganism struct {
	Name         string
	Age          uint
	Species      string
	Active       bool
	Passive      bool
	Weight       float64
	Fingerprint  []byte
	Child        Organism
	PointerChild *Organism
}

func _GetSource() ParentOrganism {

	fingerprint := sha512.Sum512([]byte("A Fingerprint"))

	return ParentOrganism{
		Name:        "Rendoru",
		Age:         22,
		Species:     "Human",
		Active:      true,
		Passive:     false,
		Weight:      172.2,
		Fingerprint: fingerprint[:],
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
		PointerChild: &Organism{
			Name:    "Ren",
			Age:     1,
			Species: "Digital",
		},
	}
}

func TestJsonStructToSameType(t *testing.T) {
	source := _GetSource()
	body, err := json.Marshal(&source)
	if err != nil {
		t.Error(err.Error())
	}
	dest := ParentOrganism{}
	err = json.Unmarshal(body, &dest)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(source, dest) {
		t.Errorf("Expected \n%v \ngot \n%v", source, dest)
	}
}

func TestMirrorStructToSameType(t *testing.T) {
	source := _GetSource()
	dest := ParentOrganism{}
	if err := Mirror(&source, &dest); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(source, dest) {
		t.Errorf("Expected \n%v \ngot \n%v", source, dest)
	}
}

func TestJsonStructToOtherType(t *testing.T) {
	source := _GetSource()
	body, err := json.Marshal(&source)
	if err != nil {
		t.Error(err.Error())
	}
	dest := AlienOrganism{}
	err = json.Unmarshal(body, &dest)
	if err != nil {
		t.Error(err.Error())
	}
	//Since the structure is different deep equal will fail
	if fmt.Sprint(source) == fmt.Sprint(dest) {
		t.Errorf("Expected \n%v \ngot \n%v", source, dest)
	}
}

func TestMirrorStructToOtherType(t *testing.T) {
	source := _GetSource()
	dest := AlienOrganism{}
	if err := Mirror(&source, &dest); err != nil {
		t.Error(err)
	}
	//Since the structure is different deep equal will fail
	if fmt.Sprint(source) != fmt.Sprint(dest) {
		t.Errorf("Expected \n%v \ngot \n%v", source, dest)
	}
}
func BenchmarkJsonStructToSameType(b *testing.B) {
	source := _GetSource()
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

func BenchmarkJsoniterStructToSameType(b *testing.B) {
	source := _GetSource()
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

func BenchmarkMirrorStructToSameType(b *testing.B) {
	source := _GetSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := ParentOrganism{}
		err := Mirror(&source, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkSmartMirrorStructToSameType(b *testing.B) {
	source := _GetSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := ParentOrganism{}
		err := SmartMirror(&source, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkJsonStructToOtherType(b *testing.B) {
	source := _GetSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		body, err := json.Marshal(&source)
		if err != nil {
			b.Error(err.Error())
		}
		dest := AlienOrganism{}
		err = json.Unmarshal(body, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkJsoniterStructToOtherType(b *testing.B) {
	source := _GetSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		body, err := jsoniter.Marshal(&source)
		if err != nil {
			b.Error(err.Error())
		}
		dest := AlienOrganism{}
		err = jsoniter.Unmarshal(body, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkMirrorStructToOtherType(b *testing.B) {
	source := _GetSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := AlienOrganism{}
		err := Mirror(&source, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkSmartMirrorStructToOtherType(b *testing.B) {
	source := _GetSource()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := AlienOrganism{}
		err := SmartMirror(&source, &dest)
		if err != nil {
			b.Error(err.Error())
		}
	}
}

func TestJsonStructToMapThenToOtherStruct(t *testing.T) {
	source := _GetSource()
	body, err := json.Marshal(&source)
	if err != nil {
		t.Error(err.Error())
	}
	destJSON := map[string]interface{}{}
	err = json.Unmarshal(body, &destJSON)
	if err != nil {
		t.Error(err.Error())
	}
	body, err = json.Marshal(destJSON)
	if err != nil {
		t.Error(err.Error())
	}
	destStruct := AlienOrganism{}
	err = json.Unmarshal(body, &destStruct)
	if err != nil {
		t.Error(err.Error())
	}
	assert.EqualValues(t, source, destStruct)
}

func TestMirrorStructToMapThenToOtherStruct(t *testing.T) {
	source := _GetSource()
	destMap := map[string]interface{}{}
	err := Mirror(&source, &destMap)
	assert.Nil(t, err)
	destStruct := AlienOrganism{}
	err = Mirror(&destMap, &destStruct)
	assert.Nil(t, err)
	assert.EqualValues(t, source, destStruct)
}

func BenchmarkJsonStructToMap(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		body, err := json.Marshal(&source)
		if err != nil {
			t.Error(err.Error())
		}
		destJSON := map[string]interface{}{}
		err = json.Unmarshal(body, &destJSON)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkJsoniterStructToMap(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		body, err := jsoniter.Marshal(&source)
		if err != nil {
			t.Error(err.Error())
		}
		destJSON := map[string]interface{}{}
		err = jsoniter.Unmarshal(body, &destJSON)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkMirrorStructToMap(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		destMap := map[string]interface{}{}
		err := Mirror(&source, &destMap)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkSmartMirrorStructToMap(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		destMap := map[string]interface{}{}
		err := SmartMirror(&source, &destMap)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkJsonMapToStruct(t *testing.B) {
	source := _GetSource()
	body, err := json.Marshal(&source)
	if err != nil {
		t.Error(err.Error())
	}
	destJSON := map[string]interface{}{}
	err = json.Unmarshal(body, &destJSON)
	if err != nil {
		t.Error(err.Error())
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		body, err = json.Marshal(destJSON)
		if err != nil {
			t.Error(err.Error())
		}
		destStruct := AlienOrganism{}
		err = json.Unmarshal(body, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkJsoniterMapToStruct(t *testing.B) {
	source := _GetSource()
	body, err := jsoniter.Marshal(&source)
	if err != nil {
		t.Error(err.Error())
	}
	destJSON := map[string]interface{}{}
	err = jsoniter.Unmarshal(body, &destJSON)
	if err != nil {
		t.Error(err.Error())
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		body, err = jsoniter.Marshal(destJSON)
		if err != nil {
			t.Error(err.Error())
		}
		destStruct := AlienOrganism{}
		err = jsoniter.Unmarshal(body, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkMirrorMapToStruct(t *testing.B) {
	source := _GetSource()
	destMap := map[string]interface{}{}
	err := Mirror(&source, &destMap)
	if err != nil {
		t.Error(err.Error())
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {

		destStruct := AlienOrganism{}
		err = Mirror(&destMap, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkSmartMirrorMapToStruct(t *testing.B) {
	source := _GetSource()
	destMap := map[string]interface{}{}
	err := SmartMirror(&source, &destMap)
	if err != nil {
		t.Error(err.Error())
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		destStruct := AlienOrganism{}
		err = SmartMirror(&destMap, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkJsonStructToMapThenToOtherStruct(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		body, err := json.Marshal(&source)
		if err != nil {
			t.Error(err.Error())
		}
		destJSON := map[string]interface{}{}
		err = json.Unmarshal(body, &destJSON)
		if err != nil {
			t.Error(err.Error())
		}
		body, err = json.Marshal(destJSON)
		if err != nil {
			t.Error(err.Error())
		}
		destStruct := AlienOrganism{}
		err = json.Unmarshal(body, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkJsoniterStructToMapThenToOtherStruct(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		body, err := jsoniter.Marshal(&source)
		if err != nil {
			t.Error(err.Error())
		}
		destJSON := map[string]interface{}{}
		err = jsoniter.Unmarshal(body, &destJSON)
		if err != nil {
			t.Error(err.Error())
		}
		body, err = jsoniter.Marshal(destJSON)
		if err != nil {
			t.Error(err.Error())
		}
		destStruct := AlienOrganism{}
		err = jsoniter.Unmarshal(body, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkMirrorStructToMapThenToOtherStruct(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		destMap := map[string]interface{}{}
		err := Mirror(&source, &destMap)
		if err != nil {
			t.Error(err.Error())
		}
		destStruct := AlienOrganism{}
		err = Mirror(&destMap, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func BenchmarkSmartMirrorStructToMapThenToOtherStruct(t *testing.B) {
	source := _GetSource()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		destMap := map[string]interface{}{}
		err := SmartMirror(&source, &destMap)
		if err != nil {
			t.Error(err.Error())
		}
		destStruct := AlienOrganism{}
		err = SmartMirror(&destMap, &destStruct)
		if err != nil {
			t.Error(err.Error())
		}
	}
}
