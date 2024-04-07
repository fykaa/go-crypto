package encrypt

func Nimbus(str string) string {
	res := ""
	for _, c := range str {
		charAscii := int(c)
		res += string(charAscii + 3)
	}
	return res
}
