package hy_go_function

import (
	"fmt"
	"hy-go-function/internal"
	"strconv"
	"testing"
)

func Test_Stream(t *testing.T) {
	array := []string{
		"1", "2", "3",
	}
	result := Of(array).Map(func(i string) int {
		atoi, err := strconv.Atoi(i)
		if err != nil {
			return 0
		}
		return atoi
	}).Filter(func(i int) bool {
		if i == 1 {
			return false
		}
		return true
	}).Map(func(i int) int {
		return i + 5
	}).Execute()

	println(fmt.Sprintf("result : %v", result))
}

func Test_filterEmpty(t *testing.T) {
	array := []string{
	}
	result := internal.DoFilter(array, func(i string) bool {
		if i == "1" {
			return false
		}
		return true
	})

	println(fmt.Sprintf("result : %v", result))
}

func Test_filter(t *testing.T) {
	array := []string{
		"1", "2", "3",
	}
	result := internal.DoFilter(array, func(i string) bool {
		if i == "1" {
			return false
		}
		return true
	})

	println(fmt.Sprintf("result : %v", result))

	result2 := internal.DoFilter(1, func(i int) bool {
		if i != 1 {
			return false
		}
		return true
	})

	println(fmt.Sprintf("result : %v", result2))
}

func Test_mapEmpty(t *testing.T) {
	array := []string{
	}
	result := internal.DoMap(array, func(i string) int {
		if i == "1" {
			return 1
		}
		return 2
	})

	println(fmt.Sprintf("result : %v", result))
}

func Test_map(t *testing.T) {
	array := []string{
		"1", "2", "3",
	}
	result := internal.DoMap(array, func(i string) int {
		if i == "1" {
			return 1
		}
		return 2
	})

	println(fmt.Sprintf("result : %v", result))

	result2 := internal.DoMap(1, func(i int) int { return i + 1 })

	println(fmt.Sprintf("result2 : %v", result2))
}
func Test_flatmap(t *testing.T) {
	array := [][]string{
		{"1", "2", "3"},
		{"1", "2", "3"},
	}
	result := internal.DoFlatMap(array, func(i []string) []string {
		return i
	})

	println(fmt.Sprintf("result : %v", result))

	result2 := internal.DoMap(1, func(i int) int { return i + 1 })

	println(fmt.Sprintf("result2 : %v", result2))
}

func Test_reduce(t *testing.T) {
	array := []int{
		1, 2, 3,
	}

	var result = internal.DoReduce(array, func(x int, y int) int {
		return x + y
	}, 0).(int)

	println(fmt.Sprintf("result : %v", result))

	var result2 = internal.DoReduce(array, func(x string, y int) string {
		return x + "," + strconv.Itoa(y)
	}, "").(string)

	println(fmt.Sprintf("result2 : %v", result2))
}
