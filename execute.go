package hy_go_function

import "hy-go-function/internal"

func (d DSL) Execute() interface{} {
	for i, operation := range d.operations {
		if d.value == nil {
			return nil
		}

		switch operation {
		case filter:
			d.value = internal.DoFilter(d.value, d.function[i])
		case maps:
			d.value = internal.DoMap(d.value, d.function[i])
		case flatMap:
			d.value = internal.DoFlatMap(d.value, d.function[i])
		}
	}
	return d.value
}
