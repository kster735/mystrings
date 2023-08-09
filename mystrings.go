package mystrings

import "fmt"

func Reverse(s string) string {
	res := ""
	for _, ch := range s {
		res = fmt.Sprintf("%c%s", ch, res)
	}
	return res
}
