package main

import (
	"fmt"
)

func arrayMerge(arr1 []string, arr2 []string) []string {
	merged := append(arr1, arr2...)

	// Membuat map untuk melacak elemen yang sudah ada
	seen := make(map[string]bool)
	result := []string{}

	for _, item := range merged {
		if _, exists := seen[item]; !exists {
			seen[item] = true
			result = append(result, fmt.Sprintf(`"%s"`, item))
		}
	}

	return result
}

func main() {

	fmt.Println(arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(arrayMerge([]string{}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(arrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(arrayMerge([]string{}, []string{}))

}
