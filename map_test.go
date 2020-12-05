package mirror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameTypeMapToMap(t *testing.T) {
	sourceStringInt := map[string]int{
		"OtherAge": 11,
		"Age":      22,
	}
	destinationStringInt := map[string]int{}
	_PerformTest("StringInt", &sourceStringInt, &destinationStringInt, false, t)

	sourceIntString := map[int]string{
		1: "This is a number",
		2: "And this is two",
		3: "This is 3",
		4: "and four",
	}
	destinationIntString := map[int]string{}
	_PerformTest("IntString", &sourceIntString, &destinationIntString, false, t)
	sourceStringString := map[string]string{
		"One":   "This is a number",
		"Two":   "And this is two",
		"Three": "This is 3",
		"Four":  "and four",
	}
	destinationStringString := map[string]string{}
	_PerformTest("StringString", &sourceStringString, &destinationStringString, false, t)
}

func TestDifferentTypeMapToMap(t *testing.T) {

	t.Run("StringInterfaceToStringString", func(t *testing.T) {
		source := map[string]interface{}{
			"OtherAge":    11,
			"Age":         uint(22),
			"Gender":      "Male",
			"RandomFloat": 1234.4567,
		}
		destination := map[string]string{}
		hasError := false
		err := SmartMirror(&source, &destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		if destination["OtherAge"] != "11" {
			t.Errorf("Expected %v got %v", "11", destination["OtherAge"])
		}

		if destination["Age"] != "22" {
			t.Errorf("Expected %v got %v", "22", destination["Age"])
		}

		if destination["Gender"] != "Male" {
			t.Errorf("Expected %v got %v", "Male", destination["Gender"])
		}

		if destination["RandomFloat"] != "1234.4567" {
			t.Errorf("Expected %v got %v", "1234.4567", destination["RandomFloat"])
		}
	})

	t.Run("StringStringToStringInt", func(t *testing.T) {
		source := map[string]string{
			"OtherAge":    "11",
			"Age":         "22",
			"Gender":      "Male",
			"RandomFloat": "1234.4567",
		}
		destination := map[string]int{}
		hasError := false
		err := SmartMirror(&source, &destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		if destination["OtherAge"] != 11 {
			t.Errorf("Expected %v got %v", "11", destination["OtherAge"])
		}

		if destination["Age"] != 22 {
			t.Errorf("Expected %v got %v", "22", destination["Age"])
		}

		if _, ok := destination["Gender"]; ok {
			t.Errorf("Expected empty but got %v", destination["Gender"])
		}

		if _, ok := destination["RandomFloat"]; ok {
			t.Errorf("Expected empty but got %v", destination["RandomFloat"])
		}
	})

}

func TestStructToMap(t *testing.T) {
	type BaseStruct struct {
		OtherAge    int
		Age         uint
		Gender      string
		RandomFloat float64
	}

	t.Run("StructToStringString", func(t *testing.T) {
		source := BaseStruct{
			OtherAge:    11,
			Age:         uint(22),
			Gender:      "Male",
			RandomFloat: 1234.4567,
		}
		destination := map[string]string{}
		hasError := false
		err := SmartMirror(&source, &destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		if destination["OtherAge"] != "11" {
			t.Errorf("Expected %v got %v", "11", destination["OtherAge"])
		}

		if destination["Age"] != "22" {
			t.Errorf("Expected %v got %v", "22", destination["Age"])
		}

		if destination["Gender"] != "Male" {
			t.Errorf("Expected %v got %v", "Male", destination["Gender"])
		}

		if destination["RandomFloat"] != "1234.4567" {
			t.Errorf("Expected %v got %v", "1234.4567", destination["RandomFloat"])
		}
	})

	t.Run("StructToStringInt", func(t *testing.T) {
		source := BaseStruct{
			OtherAge:    11,
			Age:         uint(22),
			Gender:      "Male",
			RandomFloat: 1234.4567,
		}
		destination := map[string]int{}
		hasError := false
		err := SmartMirror(&source, &destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		if destination["OtherAge"] != 11 {
			t.Errorf("Expected %v got %v", "11", destination["OtherAge"])
		}

		if destination["Age"] != 22 {
			t.Errorf("Expected %v got %v", "22", destination["Age"])
		}

		if _, ok := destination["Gender"]; ok {
			t.Errorf("Expected empty but got %v", destination["Gender"])
		}

		if destination["RandomFloat"] != 1234 {
			t.Errorf("Expected %v got %v", 1234, destination["RandomFloat"])
		}
	})

}

func TestMapToStruct(t *testing.T) {

	type Destination struct {
		A string
		B string
		C string
	}

	source := map[string]interface{}{
		"A": "aaaaa",
		"B": 12345,
		"C": "ca123",
	}

	dest := &Destination{}
	if err := SmartMirror(&source, &dest); err != nil {
		t.Error(err)
	}
	assert.Equal(t, "aaaaa", dest.A)
	assert.Equal(t, "12345", dest.B)
	assert.Equal(t, "ca123", dest.C)

}

func TestStructPointerToMap(t *testing.T) {

	type Source struct {
		A string
		B string
		C string
	}

	dest := map[string]interface{}{}

	source := &Source{
		A: "AAAA",
		B: "BBBB",
		C: "CCCC",
	}
	if err := Mirror(&source, &dest); err != nil {
		t.Error(err)
	}
	assert.Equal(t, source.A, dest["A"])
	assert.Equal(t, source.B, dest["B"])
	assert.Equal(t, source.C, dest["C"])

}
