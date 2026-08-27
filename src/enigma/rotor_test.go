package enigma

import "testing"

// TestRotorForwardBackwardAreInverses は Forward と Backward が
// 互いに逆変換になっていることを確認する。
//
// 観点: あらゆる回転位置・リングセッティングの組み合わせで、
// Forward の結果を Backward に通すと元の文字に戻ること（可逆性）。
func TestRotorForwardBackwardAreInverses(t *testing.T) {
	wiring, notch := getRotorSpec("I")

	for position := 0; position < alphabetSize; position++ {
		for ring := 0; ring < alphabetSize; ring++ {
			r := &Rotor{wiring: wiring, notch: notch, position: position, ringSetting: ring}
			for c := byte('A'); c <= 'Z'; c++ {
				got := r.Backward(r.Forward(c))
				if got != c {
					t.Fatalf("position=%d ring=%d: Backward(Forward(%q)) = %q, want %q", position, ring, c, got, c)
				}
			}
		}
	}
}

// TestRotorAtNotch は AtNotch の判定ロジックを確認する。
//
// 観点: 現在位置がノッチと一致する場合は true、一致しない場合は false を
// 正しく返すこと。
func TestRotorAtNotch(t *testing.T) {
	r := &Rotor{notch: 'V', position: charToIndex('V')}
	if !r.AtNotch() {
		t.Errorf("expected rotor at position V to be at notch V")
	}

	r.position = charToIndex('U')
	if r.AtNotch() {
		t.Errorf("expected rotor at position U not to be at notch V")
	}
}

// TestRotorStepWrapsAround は Step の境界値での挙動を確認する。
//
// 観点: 位置25（Z）から Step すると、26で割った余りにより
// 0（A）へ正しくラップアラウンドすること。
func TestRotorStepWrapsAround(t *testing.T) {
	r := &Rotor{position: alphabetSize - 1}
	r.Step()
	if r.position != 0 {
		t.Errorf("Step() from Z should wrap to A (0), got position %d", r.position)
	}
}
