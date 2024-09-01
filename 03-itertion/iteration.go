package itertion

func Repeat(letter  string) string {
	var result string

	for i := 0; i < 4; i++ {
		result +=  letter
	}

	return result
}