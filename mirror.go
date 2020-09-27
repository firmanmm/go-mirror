package mirror

import (
	"errors"
	"reflect"
)

type _RecursiveMirrorJumpTableFunc func(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error

var jumpTableRecursiveMirror map[reflect.Kind]_RecursiveMirrorJumpTableFunc

func init() {
	jumpTableRecursiveMirror = map[reflect.Kind]_RecursiveMirrorJumpTableFunc{
		reflect.Bool:      _HandleBool,
		reflect.Int:       _HandleInt,
		reflect.Int8:      _HandleInt,
		reflect.Int16:     _HandleInt,
		reflect.Int32:     _HandleInt,
		reflect.Int64:     _HandleInt,
		reflect.Uint:      _HandleUint,
		reflect.Uint8:     _HandleUint,
		reflect.Uint16:    _HandleUint,
		reflect.Uint32:    _HandleUint,
		reflect.Uint64:    _HandleUint,
		reflect.Float32:   _HandleFloat,
		reflect.Float64:   _HandleFloat,
		reflect.Slice:     _HandleList,
		reflect.Array:     _HandleList,
		reflect.String:    _HandleString,
		reflect.Map:       _HandleMap,
		reflect.Struct:    _HandleStruct,
		reflect.Interface: _HandleInterface,
		reflect.Ptr:       _HandlePointer,
	}
}

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

	sourceType := source.Type()
	destType := dest.Type()
	if sourceType == destType {
		dest.Set(source)
		return nil
	}

	if sourceKind == reflect.Interface {
		source = source.Elem()
		sourceKind = source.Kind()
	}

	if handler, ok := jumpTableRecursiveMirror[destKind]; ok {
		return handler(source, dest, sourceKind, destKind, bestEffort)
	}
	return errors.New("Destination field type didn't match Source field type")
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
//Will also attemp to convert data type to best match the destination
func SmartMirror(source, destination interface{}) error {
	return _Mirror(source, destination, true)
}
