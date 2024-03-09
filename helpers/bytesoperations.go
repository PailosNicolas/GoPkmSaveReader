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

func Uint16ToBits(value uint16) []int {
	bits := make([]int, 16)

	for i := 0; i < 16; i++ {
		bit := (value >> i) & 1
		bits[i] = int(bit)
	}

	return bits
}

func Uint32ToBits(value uint32) []int {
	bits := make([]int, 32)

	for i := 0; i < 32; i++ {
		bit := (value >> i) & 1
		bits[i] = int(bit)
	}

	return bits
}

func BitsToInt(bits []int) int {
	result := 0

	for i := len(bits) - 1; i >= 0; i-- {
		result = (result << 1) + bits[i]
	}

	return result
}
