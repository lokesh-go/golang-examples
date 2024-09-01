package main

func destCity(paths [][]string) string {
	// From city A to city B mapping
	hashMap := make(map[string]string, len(paths))
	for _, path := range paths {
		hashMap[path[0]] = path[1]
	}

	// Check which city doesn't has destination city
	for _, cityB := range hashMap {
		if _, exists := hashMap[cityB]; !exists {
			return cityB
		}
	}
	return ""
}
