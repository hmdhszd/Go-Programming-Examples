package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println("All animals", animals)
	fmt.Println("first item", animals[0])
	fmt.Println("first TWO item", animals[0:2])

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("the slice is", len(animals), "elements long")

	fmt.Println("is this slice sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("is this slice sorted now?", sort.StringsAreSorted(animals))

}
