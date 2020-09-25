package mirror

import (
	"errors"
	"reflect"
)

//Handle conversion for struct dest
func _HandleStruct(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if sourceKind == reflect.Struct {
		return _HandleStructToStruct(source, dest, bestEffort)
	} else if sourceKind == reflect.Map {
		return _HandleMapToStruct(source, dest, bestEffort)
	}
	return errors.New("Destination field type didn't match Source field type")
}

//Handle conversion from struct to struct
func _HandleStructToStruct(source, dest reflect.Value, bestEffort bool) error {
	destType := dest.Type()
	numField := destType.NumField()
	for i := 0; i < numField; i++ {
		destField := dest.Field(i)
		destName := destType.Field(i).Name
		sourceField := source.FieldByName(destName)
		if sourceField.IsValid() {
			if !(sourceField.Kind() == reflect.Ptr && sourceField.IsNil()) {
				if err := _RecursiveMirror(sourceField, destField, bestEffort); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

//Handle conversion from Map to struct
func _HandleMapToStruct(source, dest reflect.Value, bestEffort bool) error {
	destType := dest.Type()
	numField := destType.NumField()
	for i := 0; i < numField; i++ {
		destField := dest.Field(i)
		destName := destType.Field(i).Name
		sourceField := source.MapIndex(reflect.ValueOf(destName))
		if err := _RecursiveMirror(sourceField, destField, bestEffort); err != nil {
			return err
		}
	}
	return nil
}
