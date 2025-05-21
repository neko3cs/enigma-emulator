package enigma

import "strings"

type Enigma struct {
	rotors    []*Rotor
	reflector *Reflector
	plugboard *Plugboard
}

func NewEnigmaFromKey(key EnigmaKey) *Enigma {
	rotors := make([]*Rotor, 3)
	for i, name := range key.RotorOrder {
		wiring, notch := getRotorSpec(name)
		pos := int(key.RotorPositions[i] - 'A')
		ring := key.RingSettings[i] - 1
		rotors[i] = &Rotor{
			wiring:      wiring,
			notch:       notch,
			position:    pos,
			ringSetting: ring,
		}
	}
	return &Enigma{
		rotors:    rotors,
		reflector: getReflector(key.ReflectorType),
		plugboard: NewPlugboard(key.PlugboardPairs),
	}
}

func (e *Enigma) EncryptChar(c byte) byte {
	if c < 'A' || c > 'Z' {
		return c
	}

	// ステップ（ローター回転）
	if e.rotors[1].wiring[e.rotors[1].position] == e.rotors[1].notch {
		e.rotors[1].Step()
		e.rotors[2].Step()
	} else if e.rotors[0].Step() {
		e.rotors[1].Step()
	}

	c = e.plugboard.Swap(c)

	for _, rotor := range e.rotors {
		c = rotor.Forward(c)
	}

	c = e.reflector.Reflect(c)

	for i := len(e.rotors) - 1; i >= 0; i-- {
		c = e.rotors[i].Backward(c)
	}

	c = e.plugboard.Swap(c)

	return c
}

func (e *Enigma) Encrypt(text string) string {
	var result strings.Builder
	text = strings.ToUpper(text)
	for i := range len(text) {
		c := text[i]
		if c >= 'A' && c <= 'Z' {
			result.WriteByte(e.EncryptChar(c))
		}
	}
	return result.String()
}
