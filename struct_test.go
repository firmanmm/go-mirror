package mirror

import "testing"

type PrimitiveStruct struct {
	Name string
	Age  uint
}

type StructWithChild struct {
	Name  string
	Age   uint
	Child PrimitiveStruct
}

type StructWithPointerChild struct {
	Name  string
	Age   uint
	Child *PrimitiveStruct
}

type StructWithInterface struct {
	SomeAbstract interface{}
}

type StructWithInterfacePtr struct {
	SomeAbstract *interface{}
}

func TestPrimitiveStructToStruct(t *testing.T) {

	testData := []struct {
		Name        string
		Source      PrimitiveStruct
		Destination PrimitiveStruct
		HasError    bool
	}{
		{
			"StructToStruct",
			PrimitiveStruct{
				"Rendoru",
				22,
			},
			PrimitiveStruct{},
			false,
		},
	}

	for _, val := range testData {
		_PerformTest(val.Name, &val.Source, &val.Destination, val.HasError, t)
	}
}

func TestPrimitiveStructToStructWithInterface(t *testing.T) {

	testData := []struct {
		Name        string
		Source      StructWithInterface
		Destination StructWithInterfacePtr
		HasError    bool
		IsSmart     bool
	}{
		{
			"StructToStructWithEmptyInterface",
			StructWithInterface{
				SomeAbstract: nil,
			},
			StructWithInterfacePtr{},
			false,
			true,
		},
		{
			"StructToStructWithEmptyInterfaceShould without convert return err ",
			StructWithInterface{
				SomeAbstract: nil,
			},
			StructWithInterfacePtr{},
			true,
			false,
		},
	}

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			var err error
			if val.IsSmart {
				err = SmartMirror(&val.Source, &val.Destination)
			} else {
				err = Mirror(&val.Source, &val.Destination)
			}

			if err != nil {
				if !val.HasError {
					t.Errorf("Got an error %s", err.Error())
				} else {
					return
				}
			} else if err == nil && val.HasError {
				t.Error("Expecting error but got nothing")
			}
		})
	}
}

func TestStructToStructChild(t *testing.T) {

	testData := []struct {
		Name        string
		Source      StructWithChild
		Destination StructWithChild
		HasError    bool
	}{

		{
			"StructToStructChild",
			StructWithChild{
				"Rendoru",
				22,
				PrimitiveStruct{
					"Doru",
					2,
				},
			},
			StructWithChild{},
			false,
		},
	}

	for _, val := range testData {
		_PerformTest(val.Name, &val.Source, &val.Destination, val.HasError, t)
	}
}

func TestStructToStructPointerChild(t *testing.T) {

	testData := []struct {
		Name        string
		Source      StructWithPointerChild
		Destination StructWithPointerChild
		HasError    bool
	}{
		{
			"Normal",
			StructWithPointerChild{
				"Rendoru",
				22,
				&PrimitiveStruct{
					"Doru",
					2,
				},
			},
			StructWithPointerChild{},
			false,
		},
		{
			"Nil",
			StructWithPointerChild{
				"Rendoru",
				22,
				nil,
			},
			StructWithPointerChild{},
			false,
		},
	}

	for _, val := range testData {
		_PerformTest(val.Name, &val.Source, &val.Destination, val.HasError, t)
	}
}

func TestStructDifferentType(t *testing.T) {
	t.Run("Upgrade", func(t *testing.T) {
		source := PrimitiveStruct{
			Name: "Rendoru",
			Age:  22,
		}
		destination := StructWithChild{}
		hasError := false
		err := Mirror(&source, &destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		if destination.Age != source.Age {
			t.Errorf("Expected %v got %v", source.Age, destination.Age)
		}

		if destination.Name != source.Name {
			t.Errorf("Expected %v got %v", source.Name, destination.Name)
		}
	})
	t.Run("Downgrade", func(t *testing.T) {
		source := StructWithChild{
			Name: "Rendoru",
			Age:  22,
			Child: PrimitiveStruct{
				Name: "Doru",
				Age:  1,
			},
		}
		destination := PrimitiveStruct{}
		hasError := false
		err := Mirror(&source, &destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		if destination.Age != source.Age {
			t.Errorf("Expected %v got %v", source.Age, destination.Age)
		}

		if destination.Name != source.Name {
			t.Errorf("Expected %v got %v", source.Name, destination.Name)
		}
	})
}

func TestPrimitiveMapToStuct(t *testing.T) {
	source := map[string]interface{}{
		"Name": "Rendoru",
		"Age":  uint(22),
	}
	destination := PrimitiveStruct{}
	hasError := false
	err := Mirror(&source, &destination)
	if err != nil {
		if !hasError {
			t.Errorf("Got an error %s", err.Error())
		} else {
			return
		}
	} else if err == nil && hasError {
		t.Error("Expecting error but got nothing")
	}

	if destination.Age != source["Age"].(uint) {
		t.Errorf("Expected %v got %v", source["Age"], destination.Age)
	}

	if destination.Name != source["Name"] {
		t.Errorf("Expected %v got %v", source["Name"], destination.Name)
	}
}
