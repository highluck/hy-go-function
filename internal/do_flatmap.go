package internal

import (
	"reflect"
)

func DoFlatMap(input interface{}, function interface{}) interface{} {
	funcValue := reflect.ValueOf(function)
	funcType := funcValue.Type()

	inValue := reflect.ValueOf(input)
	inType := inValue.Type()
	if inType.Kind() != reflect.Slice && inType.Kind() != reflect.Array {
		result := funcValue.Call([]reflect.Value{inValue})[0]
		return result.Interface()
	}

	resultSliceType := reflect.SliceOf(funcType.Out(0).Elem())
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)
	for i := 0; i < inValue.Len(); i++ {
		elem := inValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0]
		for j := 0; j < result.Len(); j++ {
			e := result.Index(j)
			resultSlice = reflect.Append(resultSlice, e)
		}
	}

	return resultSlice.Interface()
}

