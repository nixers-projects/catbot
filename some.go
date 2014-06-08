/* 
//	missing.go - Missing golang functions for bot
//
//	authored: dami0 and dcat (at iotek dot org)
*/

package main

// check for the presence of string in string array, return -1 or index
func Present (p string, a []string) int {
	for i := 0; i < len(a); i++ {
		if (p == a[i]) { return i }
	}
	return -1
}

// remove string from string array
func Remove (pat string, lst []string) []string {
	ind := []int {}
	a := 0; b := 0; dab := 0
	d := Present(pat,lst)

	for i := d; d != -1; i += 1 + d{
		ind = append(ind, i)
		d = Present(pat,lst[i+1:])
	}
	if len(lst) == len(ind) { return []string {} }

	out := make([]string, len(lst) - len(ind))
	c := ind[0]
	if ind[0] != 0 { copy(out[:c], lst[:c]) }
	for d = 0; d < len(ind) - 1; d++ {
		a = ind[d] + 1
		b = ind[d + 1]
		dab = b - a
		c += copy(out[c:c+dab], lst[a:b])
	}
	return out
//	return ind
}
