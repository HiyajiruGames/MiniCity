package util

func IsEmpty(value *string) bool {
	return value == nil || len(*value) == 0
}
