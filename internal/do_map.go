package internal

import "reflect"

func DoMap(input interface{}, function interface{}) interface{} {
	funcValue := reflect.ValueOf(function)
	funcType := funcValue.Type()

	inValue := reflect.ValueOf(input)
	inType := inValue.Type()
	if inType.Kind() != reflect.Slice && inType.Kind() != reflect.Array {
		result := funcValue.Call([]reflect.Value{inValue})[0]
		return result.Interface()
	}

	resultSliceType := reflect.SliceOf(funcType.Out(0))
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)
	for i := 0; i < inValue.Len(); i++ {
		elem := inValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0]
		resultSlice = reflect.Append(resultSlice, result)
	}

	return resultSlice.Interface()
}
