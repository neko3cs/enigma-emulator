package enigma

// Plugboard は、信号がローターを通過する前後で、操作者が配線した
// 文字ペアを交換する部品。
type Plugboard struct {
	wiring map[byte]byte
}

// NewPlugboard は2文字のペア文字列（例: "AQ" はAとQを交換する）から
// Plugboard を組み立てる。各ペアは双方向に対称な配線となる。
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

// Swap は c に配線されている文字を返す。配線が無ければ c をそのまま返す。
func (pb *Plugboard) Swap(c byte) byte {
	if swapped, ok := pb.wiring[c]; ok {
		return swapped
	}
	return c
}
