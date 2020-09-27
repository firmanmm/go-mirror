package mirror

import "reflect"

func _HandleInterface(source, dest reflect.Value, sourceKind, destKind reflect.Kind, bestEffort bool) error {
	dest.Set(source)
	return nil
}
