package enigma

import "strings"

// Rotor は1本のエニグマローターを表す。内部配線、隣のローターを
// 回転させるノッチ位置、現在の回転位置、リングセッティングのオフセットを持つ。
type Rotor struct {
	wiring      string
	notch       byte
	position    int
	ringSetting int
}

// Forward は、プラグボード側からリフレクター側へ向かう方向で
// 文字をローターに通す。現在の回転位置とリングセッティングを考慮する。
func (r *Rotor) Forward(c byte) byte {
	index := (charToIndex(c) + r.position - r.ringSetting + alphabetSize) % alphabetSize
	wiredChar := r.wiring[index]
	output := (charToIndex(wiredChar) - r.position + r.ringSetting + alphabetSize) % alphabetSize
	return indexToChar(output)
}

// Backward は、リフレクター側からプラグボード側へ向かう方向で
// 文字をローターに通す。Forward の逆変換にあたる。
func (r *Rotor) Backward(c byte) byte {
	index := (charToIndex(c) + r.position - r.ringSetting + alphabetSize) % alphabetSize
	wiredIndex := strings.IndexByte(r.wiring, indexToChar(index))
	output := (wiredIndex - r.position + r.ringSetting + alphabetSize) % alphabetSize
	return indexToChar(output)
}

// AtNotch は、ローターが現在ノッチの位置にあるかどうかを返す。
// true のとき、次の打鍵でこのローターの左隣のローターも回転する
// （Enigma.stepRotors を参照）。
func (r *Rotor) AtNotch() bool {
	return indexToChar(r.position) == r.notch
}

// Step はローターの位置を1つ進める。Z の次は A に戻る（ラップアラウンド）。
func (r *Rotor) Step() {
	r.position = (r.position + 1) % alphabetSize
}

// getRotorSpec は、実機のローター名（"I"〜"V"）に対応する配線表と
// ノッチ文字を返す。それ以外の名前を渡すと panic する。
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
