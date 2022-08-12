package iteration

func Repeat(character string, repeatNo int) (repeated string) {
	for i := 0; i < repeatNo; i++ {
		repeated += character
	}
	return
}
