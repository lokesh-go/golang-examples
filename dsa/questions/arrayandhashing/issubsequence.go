package main

func isSubsequence(s string, t string) bool {
	i := 0
	for ind := range t {
		if i < len(s) && t[ind] == s[i] {
			i++
		}
	}

	if i == len(s) {
		return true
	}
	return false
}
