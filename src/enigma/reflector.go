package enigma

type Reflector struct {
	wiring string
}

func (r *Reflector) Reflect(c byte) byte {
	index := c - 'A'
	return r.wiring[index]
}

func getReflector(name string) *Reflector {
	switch name {
	case "B":
		return &Reflector{wiring: "YRUHQSLDPXNGOKMIEBFZCWVJAT"}
	case "C":
		return &Reflector{wiring: "FVPJIAOYEDRZXWGCTKUQSBNMHL"}
	default:
		panic("Invalid reflector: " + name)
	}
}
