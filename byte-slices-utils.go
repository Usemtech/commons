package commons

// LengthOfByteSlices should be moved to `utils`
func LengthOfByteSlices(slices ...[]byte) (length int) {
	for _, s := range slices {
		length += len(s)
	}

	return
}

// ConcatByteSlices should be moved to `utils`
func ConcatByteSlices(slices ...[]byte) (bytes []byte) {
	length := LengthOfByteSlices(slices...)
	bytes = make([]byte, length)
	var offset int
	for _, s := range slices {
		offset += copy(bytes[offset:], s)
	}

	return
}
