package mirror

import (
	"errors"
	"reflect"
)

//Handle conversion for map dest
func _HandleMap(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if sourceKind == reflect.Map {
		return _HandleMapToMap(source, dest, bestEffort)
	} else if sourceKind == reflect.Struct {
		return _HandleStructToMap(source, dest, bestEffort)
	}
	return errors.New("Destination field type didn't match Source field type")
}

//Handle conversion for map to map
//Basically copy source to destination but allow only correct key and value
func _HandleMapToMap(source, dest reflect.Value, bestEffort bool) error {
	mapEntry := source.MapRange()
	destType := dest.Type()
	destKey := destType.Key()
	destValue := destType.Elem()
	for mapEntry.Next() {
		key := reflect.New(destKey).Elem()
		value := reflect.New(destValue).Elem()
		if err := _RecursiveMirror(mapEntry.Key(), key, bestEffort); err != nil {
			if bestEffort {
				continue
			}
			return err
		}
		if key.IsZero() {
			continue
		}
		if err := _RecursiveMirror(mapEntry.Value(), value, bestEffort); err != nil {
			if bestEffort {
				continue
			}
			return err
		}
		if value.IsZero() {
			continue
		}
		dest.SetMapIndex(key, value)
	}
	return nil
}

//Handle conversion for struct to map
//Basically copy struct field to destination but allow only correct key and value
func _HandleStructToMap(source, dest reflect.Value, bestEffort bool) error {
	sourceType := source.Type()
	numField := source.NumField()
	destType := dest.Type()
	destKey := destType.Key()
	destValue := destType.Elem()
	for i := 0; i < numField; i++ {
		sourceField := source.Field(i)
		sourceName := sourceType.Field(i).Name

		key := reflect.New(destKey).Elem()
		value := reflect.New(destValue).Elem()
		if err := _RecursiveMirror(reflect.ValueOf(sourceName), key, bestEffort); err != nil {
			if bestEffort {
				continue
			}
			return err
		}
		if key.IsZero() {
			continue
		}
		if err := _RecursiveMirror(sourceField, value, bestEffort); err != nil {
			if bestEffort {
				continue
			}
			return err
		}
		if value.IsZero() {
			continue
		}
		dest.SetMapIndex(key, value)

	}
	return nil
}
