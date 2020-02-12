package iteration

// Repeat returns repeated string given specific count
func Repeat(s string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += s
	}
	return
}
