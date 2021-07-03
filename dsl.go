package hy_go_function

type operation string

const (
	filter  = operation("filter")
	maps    = operation("map")
	flatMap = operation("flatmap")
)

type DSL struct {
	function      []interface{}
	value         interface{}
	operations    []operation
}

func Of(value interface{}) DSL {
	return DSL{
		value: value,
	}
}

func (d DSL) Filter(predicate interface{}) DSL {
	d.operations = append(d.operations, filter)
	d.function = append(d.function, predicate)
	return d
}

func (d DSL) FlatMap(function interface{}) DSL {
	d.operations = append(d.operations, flatMap)
	d.function = append(d.function, function)
	return d
}

func (d DSL) Map(function interface{}) DSL {
	d.operations = append(d.operations, maps)
	d.function = append(d.function, function)
	return d
}