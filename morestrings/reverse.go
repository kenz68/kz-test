package morestrings

// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "string" package.

// ReverseRuns returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	lth := len(r)
	for i := 0; i < lth/2; i++ {
		j := lth - 1 - i
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
