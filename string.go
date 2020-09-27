package mirror

import (
	"errors"
	"fmt"
	"reflect"
)

func _HandleString(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if !bestEffort {
		if destKind != sourceKind {
			return errors.New("Destination field type didn't match Source field type")
		}
		dest.Set(source)
	} else {
		switch sourceKind {
		case reflect.String:
			dest.Set(source)
		case reflect.Int, reflect.Uint, reflect.Float32, reflect.Float64, reflect.Bool:
			dest.SetString(fmt.Sprint(source.Interface()))
		default:
			return errors.New("Destination field type didn't match Source field type")
		}
	}
	return nil
}
