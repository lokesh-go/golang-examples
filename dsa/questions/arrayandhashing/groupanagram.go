package main

/*
Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	1. Make a hashMap for each word so that KEY will contains the same values
		[1 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0]  --  [eat tea ate]
		[1 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 1 0 0 0 0 0 0]  --  [tan nat]
		[1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0]  --  [bat]

	2. For Key building
		type key [26]byte
		var k key
		k[str[i]-'a']++ // For each word Range over str

	3. Build Hash Map
		var hashMap map[key][]string
		hashMap[k] = append(hashMap[k], str)
*/

func groupAnagrams(input []string) [][]string {
	type key [26]byte
	g := make(map[key][]string, len(input))
	for _, str := range input {
		k := key{}
		for i := range str {
			k[str[i]-'a']++
		}
		g[k] = append(g[k], str)
	}

	res := make([][]string, 0, len(g))
	for _, ga := range g {
		res = append(res, ga)
	}

	return res
}

/*
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
*/
