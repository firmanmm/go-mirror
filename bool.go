package mirror

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func _HandleBool(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if !bestEffort {
		if destKind != sourceKind {
			return errors.New("Destination field type didn't match Source field type")
		}
		dest.Set(source)
	} else {
		switch sourceKind {
		case reflect.Bool:
			dest.Set(source)
		case reflect.Int:
			intVal := source.Int()
			dest.SetBool(intVal > 0)
		case reflect.Uint:
			uintVal := source.Uint()
			dest.SetBool(uintVal > 0)
		case reflect.String:
			rawString := source.String()
			val, err := strconv.ParseBool(rawString)
			if err != nil {
				return fmt.Errorf("Failed to mirror String to Bool, err : %s", err.Error())
			}
			dest.SetBool(val)
		default:
			return errors.New("Destination field type didn't match Source field type")
		}
	}
	return nil
}
