package main

func main() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// print all values of a map
	for key, value := range intMap {
		println(key, value)
	}

	// delete an item of a map
	delete(intMap, "four")

	//check if an item exists in a map
	el, ok := intMap["two"]

	if ok {
		println(el, "is in map")
	} else {
		println(el, "is not in map")
	}
}