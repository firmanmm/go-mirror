package mirror

import (
	"reflect"
)

//Handle conversion for list dest
//Will add all element from source to dest
func _HandleList(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	length := source.Len()
	destType := dest.Type()
	destValue := destType.Elem()
	for i := 0; i < length; i++ {
		value := reflect.New(destValue).Elem()
		if err := _RecursiveMirror(source.Index(i), value, bestEffort); err != nil {
			if bestEffort {
				continue
			}
			return err
		}
		dest.Set(reflect.Append(dest, value))
	}
	return nil
}
