package mirror

import (
	"reflect"
)

func _HandlePointer(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {

	sourceType := source.Type()
	destType := dest.Type()
	if sourceType == destType {
		dest.Set(source)
		return nil
	}
	if sourceKind == reflect.Ptr {
		source = source.Elem()
	}
	newDest := reflect.New(dest.Type())
	dest.Set(newDest.Elem())
	return _RecursiveMirror(source, dest, bestEffort)
}
