package enigma

import "strings"

// Enigma は、ローター群・リフレクター・プラグボードを組み合わせた、
// 設定済みのエニグマ機。文字列の暗号化（対称性により復号にも使える）を行う。
type Enigma struct {
	rotors    []*Rotor
	reflector *Reflector
	plugboard *Plugboard
}

// NewEnigmaFromKey は key からエニグマ機を組み立てる。ローター名・
// リフレクター名を配線表に解決し、プラグボードの配線も設定する。
func NewEnigmaFromKey(key EnigmaKey) *Enigma {
	rotors := make([]*Rotor, 3)
	for i, name := range key.RotorOrder {
		wiring, notch := getRotorSpec(name)
		pos := charToIndex(key.RotorPositions[i])
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

// EncryptChar は英大文字1文字を機械に通す。ローターを回転させたのち、
// プラグボード→ローター群（順方向）→リフレクター→ローター群（逆方向）
// →プラグボードの順で信号を通す。A〜Z以外の文字はそのまま返す。
func (e *Enigma) EncryptChar(c byte) byte {
	if c < 'A' || c > 'Z' {
		return c
	}

	e.stepRotors()

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

// stepRotors は打鍵1回分、ローターを回転させる。実機のダブルステッピング
// 挙動を再現しており、右（高速）ローターは毎打鍵無条件に回転する。
// 中央ローターは、自身がノッチにある場合（このとき左のローターも回転する）、
// または今回の打鍵の前に右ローターがノッチにあった場合に、追加で回転する。
func (e *Enigma) stepRotors() {
	fast, middle, slow := e.rotors[0], e.rotors[1], e.rotors[2]
	if middle.AtNotch() {
		slow.Step()
		middle.Step()
	} else if fast.AtNotch() {
		middle.Step()
	}
	fast.Step()
}

// Encrypt は text を大文字化し、英字だけを EncryptChar に通す。
// 英字以外の文字は結果から取り除かれる。同じ鍵で組み立てた別の機械に
// 出力を再度通すと、元の文字列に戻る。
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
