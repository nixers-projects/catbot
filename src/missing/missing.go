/*
//   missing.go - missing golang functions
//
//   authors: dcat, dami0
*/
package missing

//import "fmt"

// check for presence of string in string array
func Present (t []string, q string) bool {
        for i :=0; i < len(t); i++ {
                if (q == t[i]) { return true }
        }
        return false
}

// remove string from string array
func Remove (string_list []string, pattern string) []string {
        count := len(string_list)
        for i := 0; i < count; i++ {
                if (string_list[i] == pattern) {
                        temp := make([]string, count - 1)
                        copy(temp, string_list[:i])
                        copy(temp[i:], string_list[i+1:])
                        string_list = temp
                        count -= 1
                        i -= 1
                }
        }
        
        return string_list
}
