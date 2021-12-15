package mirror

import "testing"

type MyString string

func TestPrimitiveStringToString(t *testing.T) {

	testData := []struct {
		Name        string
		Source      string
		Destination string
		HasError    bool
	}{
		{
			"StringToString",
			"test",
			"",
			false,
		},
	}

	for _, val := range testData {
		_PerformTest(val.Name, &val.Source, &val.Destination, val.HasError, t)
	}
}

func TestPrimitiveStringToMyString(t *testing.T) {

	testData := []struct {
		Name        string
		Source      string
		Destination MyString
		Expect      MyString
		HasError    bool
	}{
		{
			"StringToMyString",
			"test",
			"",
			"test",
			false,
		},
	}

	for _, val := range testData {
		_PerformTestSmart(val.Name, &val.Source, &val.Destination, &val.Expect, val.HasError, t)
	}
}

func TestPrimitiveMyStringToString(t *testing.T) {

	testData := []struct {
		Name        string
		Source      MyString
		Destination string
		Expect      string
		HasError    bool
	}{
		{
			"MyStringToString",
			"test",
			"",
			"test",
			false,
		},
	}

	for _, val := range testData {
		_PerformTestSmart(val.Name, &val.Source, &val.Destination, &val.Expect, val.HasError, t)
	}
}
