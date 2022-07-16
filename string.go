package string

type String struct {
}

// Reverse returns the reversed string.
// The idea is swapping first index and last index of each round by looping.
func (s *String) Reverse(text []byte) string {
	for i := 0; i < len(text)/2; i++ {
		text[i], text[len(text)-1-i] = text[len(text)-1-i], text[i]
	}

	return string(text)
}
