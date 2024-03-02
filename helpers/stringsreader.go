package helpers

import "strings"

func ReadString(text []byte) string {
	chars := "0123456789!?.-         ,  ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := ""
	for _, i := range text {
		c := int(i) - 161
		if c < 0 || c >= len(chars) {
			ret = ret + " "
		} else {
			ret = ret + string(chars[c])
		}
	}
	return strings.TrimSpace(ret)
}
