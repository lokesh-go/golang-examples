package main

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
