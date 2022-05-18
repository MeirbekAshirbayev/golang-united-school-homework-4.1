package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	s1 := ""
	m := []rune{}
	for _, v := range input {
		m = append(m, v)
	}

	for i := range m {
		s1 = s1 + string(m[len(m)-1-i])
	}
	output = s1
	return output
}
