package capitalize

// Returns a copy of string with
// the first character converted to uppercase and
// the remainder to lowercase.
func Str(s string) string {
	if len(s) == 0 {
		return ""
	}
	tmp := []rune(s)
	for i := 0; i < len(tmp); i++ {
		if i == 0 {
			if tmp[i] > 96 && tmp[i] < 123 {
				tmp[i] -= 32
			}
		} else {
			if tmp[i] > 64 && tmp[i] < 91 {
				tmp[i] += 32
			}
		}
	}
	return string(tmp)
}
