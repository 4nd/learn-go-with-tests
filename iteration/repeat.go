package iteration

func Repeat(char string, repeatCount int) string {
	var repeated string
	for range repeatCount {
		repeated += char
	}
	return repeated
}
