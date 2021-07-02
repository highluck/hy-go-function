# hy-go-function
go-functional
``` go
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
  ``` 
