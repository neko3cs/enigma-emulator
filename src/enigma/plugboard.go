package enigma

type Plugboard struct {
	wiring map[byte]byte
}

func NewPlugboard(pairs []string) *Plugboard {
	pb := &Plugboard{wiring: make(map[byte]byte)}
	for _, pair := range pairs {
		a := pair[0]
		b := pair[1]
		pb.wiring[a] = b
		pb.wiring[b] = a
	}
	return pb
}

func (pb *Plugboard) Swap(c byte) byte {
	if swapped, ok := pb.wiring[c]; ok {
		return swapped
	}
	return c
}
