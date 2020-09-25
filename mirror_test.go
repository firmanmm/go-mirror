package mirror

import (
	"reflect"
	"testing"
)

func _PerformTest(name string, source, destination interface{}, hasError bool, t *testing.T) {
	t.Run(name, func(t *testing.T) {
		err := Mirror(source, destination)
		if err != nil {
			if !hasError {
				t.Errorf("Got an error %s", err.Error())
			} else {
				return
			}
		} else if err == nil && hasError {
			t.Error("Expecting error but got nothing")
		}

		source = reflect.ValueOf(source).Elem().Interface()
		destination = reflect.ValueOf(destination).Elem().Interface()
		if !reflect.DeepEqual(source, destination) {
			t.Errorf("Expected %v but got %v", source, destination)
		}
	})
}

func TestPrimitive(t *testing.T) {

	inInt := -100
	outInt := 0
	_PerformTest("Int", &inInt, &outInt, false, t)

	inUint := -100
	outUint := 0
	_PerformTest("Uint", &inUint, &outUint, false, t)

	inFloat := float64(1234.5678)
	outFloat := float64(0)
	_PerformTest("Float", &inFloat, &outFloat, false, t)

	inString := "This is a string"
	outString := ""
	_PerformTest("String", &inString, &outString, false, t)

	inBool := true
	outBool := false
	_PerformTest("Bool", &inBool, &outBool, false, t)

	outUint = 0
	_PerformTest("CrossIntUint", &inInt, &outUint, false, t)

	outFloat = 0
	_PerformTest("CrossIntFloat", &inInt, &outFloat, true, t)

	outString = ""
	_PerformTest("CrossIntString", &inInt, &outString, true, t)

}
