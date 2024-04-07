package decrypt

func Nimbus(str string) string {
	res := ""
	for _, char := range str {
		asciiCode := int(char)
		res += string(asciiCode - 3)
	}
	return res
}
