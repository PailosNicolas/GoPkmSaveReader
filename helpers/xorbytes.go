package helpers

func XorBytes(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("byte slices must have equal length")
	}
	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}
