package util

func GetReverseMap(src map[int64]string) map[string]int64 {
	data := make(map[string]int64)
	for key, value := range src {
		data[value] = key
	}
	return data
}
