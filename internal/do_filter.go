package internal

import "reflect"

func DoFilter(array interface{}, predicate interface{}) interface{} {
	funcValue := reflect.ValueOf(predicate)
	funcType := funcValue.Type()

	if funcType.Out(0).Kind() != reflect.Bool {
		panic("type not boolean")
	}

	inValue := reflect.ValueOf(array)
	inType := inValue.Type()

	if inType.Kind() != reflect.Slice && inType.Kind() != reflect.Array {
		result := funcValue.Call([]reflect.Value{inValue})[0].Interface().(bool)
		if result {
			return inValue
		} else {
			return nil
		}
	}

	resultSliceType := reflect.SliceOf(inType.Elem())
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

	for i := 0; i < inValue.Len(); i++ {
		elem := inValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)
		if result {
			resultSlice = reflect.Append(resultSlice, elem)
		}
	}

	return resultSlice.Interface()
}
