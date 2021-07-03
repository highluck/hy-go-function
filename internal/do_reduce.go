package internal

import "reflect"

func DoReduce(input interface{}, function interface{}, initial interface{}) interface{} {
	funcValue := reflect.ValueOf(function)

	inValue := reflect.ValueOf(input)
	inType := inValue.Type()
	if inType.Kind() != reflect.Slice && inType.Kind() != reflect.Array {
		result := funcValue.Call([]reflect.Value{inValue})[0]
		return result.Interface()
	}

	initialValue := reflect.ValueOf(initial)
	for i := 0; i < inValue.Len(); i++ {
		elem := inValue.Index(i)
		initialValue = funcValue.Call([]reflect.Value{initialValue, elem})[0]
	}
	return initialValue.Interface()
}
