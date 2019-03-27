package redis

func argsForm(list []interface{}, args ...interface{}) []interface{} {
	arr := make([]interface{}, 0)
	arr = append(arr, args...)
	arr = append(arr, list...)
	return arr
}
