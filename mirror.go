package mirror

import (
	"errors"
	"reflect"
)

func _RecursiveMirror(source, dest reflect.Value, bestEffort bool) error {

	destKind := dest.Kind()
	sourceKind := source.Kind()

	if !dest.CanSet() {
		return nil
	}

	switch sourceKind {
	case reflect.Ptr, reflect.Slice, reflect.Map:
		if source.IsNil() {
			return nil
		}
	default:
	}
	if sourceKind == reflect.Interface {
		source = source.Elem()
		sourceKind = source.Kind()
	}

	switch destKind {
	case reflect.Struct:
		return _HandleStruct(source, dest, sourceKind, destKind, bestEffort)
	case reflect.Map:
		return _HandleMap(source, dest, sourceKind, destKind, bestEffort)
	case reflect.Slice:
		return _HandleList(source, dest, sourceKind, destKind, bestEffort)
	case reflect.Int:
		return _HandleInt(source, dest, sourceKind, destKind, bestEffort)
	case reflect.Uint:
		return _HandleUint(source, dest, sourceKind, destKind, bestEffort)
	case reflect.Float32:
		return _HandleFloat(source, dest, sourceKind, destKind, bestEffort)
	case reflect.String:
		return _HandleString(source, dest, sourceKind, destKind, bestEffort)
	case sourceKind:
		dest.Set(source)
	default:
		if !bestEffort {
			return errors.New("Destination field type didn't match Source field type")
		}
	}
	return nil
}

func _Mirror(source, destination interface{}, bestEffort bool) error {
	src := reflect.ValueOf(source)
	dest := reflect.ValueOf(destination)
	if dest.Kind() == reflect.Ptr {
		dest = dest.Elem()
	}
	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}
	if !dest.CanSet() {
		return errors.New("Destination is not set-able, are you passing non pointer value?")
	}
	return _RecursiveMirror(src, dest, bestEffort)
}

//Convert arbitrary interface to certain structure
//Will NOT attemp to convert data type, see [SmartMirror]
func Mirror(source, destination interface{}) error {
	return _Mirror(source, destination, false)
}

//Convert arbitrary interface to certain structure
//but will not give error when conversion is failed.
//Will also attemp to convert data type to best match the destination
func SmartMirror(source, destination interface{}) error {
	return _Mirror(source, destination, true)
}
