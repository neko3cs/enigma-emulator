package enigma

import "strings"

type Rotor struct {
	wiring      string
	notch       byte
	position    int
	ringSetting int
}

func (r *Rotor) Forward(c byte) byte {
	index := (charToIndex(c) + r.position - r.ringSetting + alphabetSize) % alphabetSize
	wiredChar := r.wiring[index]
	output := (charToIndex(wiredChar) - r.position + r.ringSetting + alphabetSize) % alphabetSize
	return indexToChar(output)
}

func (r *Rotor) Backward(c byte) byte {
	index := (charToIndex(c) + r.position - r.ringSetting + alphabetSize) % alphabetSize
	wiredIndex := strings.IndexByte(r.wiring, indexToChar(index))
	output := (wiredIndex - r.position + r.ringSetting + alphabetSize) % alphabetSize
	return indexToChar(output)
}

func (r *Rotor) AtNotch() bool {
	return indexToChar(r.position) == r.notch
}

func (r *Rotor) Step() {
	r.position = (r.position + 1) % alphabetSize
}

func getRotorSpec(name string) (string, byte) {
	switch name {
	case "I":
		return "EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'Q'
	case "II":
		return "AJDKSIRUXBLHWTMCQGZNPYFVOE", 'E'
	case "III":
		return "BDFHJLCPRTXVZNYEIWGAKMUSQO", 'V'
	case "IV":
		return "ESOVPZJAYQUIRHXLNFTGKDCMWB", 'J'
	case "V":
		return "VZBRGITYUPSDNHLXAWMJQOFECK", 'Z'
	default:
		panic("Invalid rotor name: " + name)
	}
}
