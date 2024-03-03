package helpers

func DecryptData(encryptedData []byte, decryptionKey []byte) []byte {
	decryptedData := make([]byte, len(encryptedData))
	for i := range 3 {
		start := i * 4
		encryptedChunk := encryptedData[start : start+4]
		decryptedChunk := XorBytes(encryptedChunk, decryptionKey)
		copy(decryptedData[start:], decryptedChunk)
	}
	return decryptedData
}
