package enigma

// Reflector は、信号を折り返して逆方向にローター群へ通す部品。
// その配線は「対合（involution）」になっている（2回反射すると元の文字に
// 戻る）。これがエニグマの暗号化を、暗号化と復号の両方に使える性質
// （対称性）の根拠になっている。
type Reflector struct {
	wiring string
}

// Reflect は c がリフレクター上で配線されている文字を返す。
func (r *Reflector) Reflect(c byte) byte {
	return r.wiring[charToIndex(c)]
}

// getReflector は、実機のリフレクター名（"B" または "C"）に対応する
// 配線を返す。それ以外の名前を渡すと panic する。
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
