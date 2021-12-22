package main

import "fmt"

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

	//---------------------------------------------------------

	MyMap := map[string]int{
		"Hamid":           30,
		"James":           32,
		"Miss MoneyPenny": 27,
	}

	fmt.Println(MyMap)
	fmt.Println(MyMap["Hamid"])
	fmt.Println(MyMap["NoBody"])

	//---------------------------------------------------------

	NewMap := map[string][]string{
		`James Bond`:         []string{`First item of James Bond`, `Second Item of James Bond`, `Third Item of James Bond`},
		`Hamid Hosseinzadeh`: []string{`First item of Hamid Hosseinzadeh`, `Second Item of Hamid Hosseinzadeh`, `Third Item of Hamid Hosseinzadeh`},
		`Harry Potter`:       []string{`First item of Harry Potter`, `Second Item of Harry Potter`, `Third Item of Harry Potter`},
	}

	NewMap["New-Item-Of-The-New-Map"] = []string{"1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th"}

	for item, value := range NewMap {

		fmt.Printf("\n Item: %v \n \t Value:\n", item)

		for index, val := range value {
			fmt.Printf("\t \t %v- %v \n", index, val)
		}

	}

}
