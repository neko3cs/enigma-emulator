package enigma

import "strings"

type Rotor struct {
	wiring      string
	notch       byte
	position    int
	ringSetting int
}

func (r *Rotor) Forward(c byte) byte {
	index := (int(c-'A') + r.position - r.ringSetting + 26) % 26
	wiredChar := r.wiring[index]
	output := (int(wiredChar-'A') - r.position + r.ringSetting + 26) % 26
	return byte(output) + 'A'
}

func (r *Rotor) Backward(c byte) byte {
	index := (int(c-'A') + r.position - r.ringSetting + 26) % 26
	wiredIndex := strings.IndexByte(r.wiring, byte(index)+'A')
	output := (wiredIndex - r.position + r.ringSetting + 26) % 26
	return byte(output) + 'A'
}

func (r *Rotor) Step() bool {
	r.position = (r.position + 1) % 26
	return r.notch == r.wiring[r.position]
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
