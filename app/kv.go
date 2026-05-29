package main

var kvStore = make(map[string]string)

func setKeyValue(key string, value string) {
	kvStore[key] = value
}

func getKeyValue(key string) (string, bool) {
	value, exists := kvStore[key]
	return value, exists
}