package hArrays

import "reflect"

//InArray check if value in array
func InArray(array interface{}, value interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		sliceLength := s.Len()
		for i := 0; i < sliceLength; i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}
