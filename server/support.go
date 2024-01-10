package main

func Contains[T comparable](arr []T, t T) bool {
	for _, el := range arr {
		if el == t {
			return true
		}
	}
	return false
}
