package mirror

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func _HandleInt(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if !bestEffort {
		if destKind != sourceKind {
			return errors.New("Destination field type didn't match Source field type")
		}
		dest.Set(source)
	} else {

		switch sourceKind {
		case reflect.Int:
			dest.Set(source)
		case reflect.Uint:
			dest.SetInt(int64(source.Interface().(uint)))
		case reflect.Float32:
			dest.SetInt(int64(source.Interface().(float32)))
		case reflect.Float64:
			dest.SetInt(int64(source.Interface().(float64)))
		case reflect.String:
			rawString := source.Interface().(string)
			number, err := strconv.ParseInt(rawString, 10, 0)
			if err != nil {
				if !bestEffort {
					return fmt.Errorf("Failed to mirror String to Int, err : %s", err.Error())
				}
				return nil
			}
			dest.SetInt(number)
		default:
			if bestEffort {
				return nil
			}
			return errors.New("Destination field type didn't match Source field type")
		}
	}
	return nil
}

func _HandleUint(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if !bestEffort {
		if destKind != sourceKind {
			return errors.New("Destination field type didn't match Source field type")
		}
		dest.Set(source)
	} else {
		switch sourceKind {
		case reflect.Uint:
			dest.Set(source)
		case reflect.Int:
			dest.SetUint(uint64(source.Interface().(int)))
		case reflect.Float32:
			dest.SetUint(uint64(source.Interface().(float32)))
		case reflect.Float64:
			dest.SetUint(uint64(source.Interface().(float64)))
		case reflect.String:
			rawString := source.Interface().(string)
			number, err := strconv.ParseUint(rawString, 10, 0)
			if err != nil {
				if !bestEffort {
					return fmt.Errorf("Failed to mirror String to Int, err : %s", err.Error())
				}
				return nil
			}
			dest.SetUint(number)
		default:
			if bestEffort {
				return nil
			}
			return errors.New("Destination field type didn't match Source field type")
		}
	}
	return nil
}

func _HandleFloat(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	if !bestEffort {
		if destKind != sourceKind {
			return errors.New("Destination field type didn't match Source field type")
		}
		dest.Set(source)
	} else {

		switch sourceKind {
		case reflect.Float32, reflect.Float64:
			dest.Set(source)
		case reflect.Int:
			dest.SetFloat(float64(source.Interface().(int)))
		case reflect.Uint:
			dest.SetFloat(float64(source.Interface().(uint)))
		case reflect.String:
			rawString := source.Interface().(string)
			number, err := strconv.ParseFloat(rawString, 0)
			if err != nil {
				if !bestEffort {
					return fmt.Errorf("Failed to mirror String to Int, err : %s", err.Error())
				}
				return nil
			}
			dest.SetFloat(number)
		default:
			if bestEffort {
				return nil
			}
			return errors.New("Destination field type didn't match Source field type")
		}
	}
	return nil
}
