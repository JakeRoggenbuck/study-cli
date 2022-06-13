package main

func indexOf(word string, data [4]string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
