/* 
//	missing.go - Missing golang functions for bot
//
//	authored: dami0 and dcat (at iotek dot org)
*/

package some

// check for the presence of string in string array, return -1 or index
func Present (p string, a []string) int {
	for i := 0; i < len(a); i++ {
		if (p == a[i]) { return i }
	}
	return -1
}

// remove (all instances) string from string array
func Remove (pat string, lst []string) []string {
	ind := []int {}
	a := 0; b := 0; dab := 0
	d := Present(pat, lst)

	for i := d; d != -1; i += 1 + d {
		ind = append(ind, i)
		d = Present(pat,lst[i+1:])
	}
	ret_l := len(lst) - len(ind)

	if ret_l == len(lst) { return lst }
	c := ind[0]
	if ret_l == 0 {
		return []string {}
	} else if ret_l == 1 {
		return []string {}
	} else if len(ind) == 1 {
		copy(lst[:c], lst[:c])
		copy(lst[c:], lst[c+1:])
		return lst[:ret_l]
	}

	out := lst
	for d = 0; d < len(ind) - 1; d++ {
		a = ind[d] + 1
		b = ind[d + 1]
		dab = b - a
		if dab > 0 { c += copy(out[c:c+dab], lst[a:b]) }
	}
	copy(out[c:], lst[ind[d] + 1:])
	return out[:ret_l]
}
