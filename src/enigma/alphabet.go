package enigma

const alphabetSize = 26

func charToIndex(c byte) int {
	return int(c - 'A')
}

func indexToChar(i int) byte {
	return byte(i) + 'A'
}
