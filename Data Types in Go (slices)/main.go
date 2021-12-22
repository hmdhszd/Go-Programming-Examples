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

	// ------------------------------------------------------------------

	// MyNewSlice := make(type, length, capacity)
	MyNewSlice := make([]int, 10, 100)
	fmt.Println(MyNewSlice)
	fmt.Println(len(MyNewSlice))
	fmt.Println(cap(MyNewSlice))

	MyNewSlice = append(MyNewSlice, 1234, 21345, 2346, 2345)
	fmt.Println(MyNewSlice)
	fmt.Println(len(MyNewSlice))
	fmt.Println(cap(MyNewSlice))

	// ------------------------------------------------------------------
	// Multi dimentional slice

	s1 := []string{"Hamid", "Hosseinzadeh", "1991", "Jan", "24"}
	s2 := []string{"James", "Bond", "1982", "Mai", "02"}

	s12 := [][]string{s1, s2}

	fmt.Println(s12)

}
