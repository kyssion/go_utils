package util

func XOR(src bool, target bool) bool {
	return (src || target) && !(src && target)
}

func GetPtrBool(val bool) *bool {
	return &val
}
