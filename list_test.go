package mirror

import (
	"testing"
)

func TestSameTypeList(t *testing.T) {
	sourceString := []string{"A", "B", "C"}
	destinationString := []string{}
	_PerformTest("String", &sourceString, &destinationString, false, t)

	sourceInt := []int{1, 2, 3}
	destinationInt := []int{}
	_PerformTest("Int", &sourceInt, &destinationInt, false, t)
}
