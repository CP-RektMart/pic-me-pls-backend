package convert

import "reflect"

// return nil if data is default value
func ToPointer[T any](data T) *T {
	var t T
	if reflect.DeepEqual(t, data) {
		return nil
	}
	return &data
}
