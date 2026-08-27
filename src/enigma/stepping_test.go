package enigma

import "testing"

// 標準的なエニグマのダブルステッピング挙動を回帰させないための固定シーケンステスト。
// rotors III(fast)/II(middle)/I(slow), 位置AAAから100回ステップさせた結果は
// 既知の正しいアルゴリズムでシミュレートした WEA と一致するはずである。
func TestStepRotorsDoubleStepping(t *testing.T) {
	e := NewEnigmaFromKey(EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "AAA",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: nil,
		ReflectorType:  "B",
	})

	for i := 0; i < 100; i++ {
		e.stepRotors()
	}

	got := string([]byte{
		indexToChar(e.rotors[0].position),
		indexToChar(e.rotors[1].position),
		indexToChar(e.rotors[2].position),
	})
	want := "WEA"
	if got != want {
		t.Errorf("after 100 steps, rotor positions = %q, want %q", got, want)
	}
}

// 右（高速）ローターは、中央ローターが自身のノッチにある打鍵（ダブルステップ）でも
// 必ず回転しなければならない。過去にここが抜けていた（AGENTS.md Incidents参照）。
func TestFastRotorAlwaysSteps(t *testing.T) {
	fast := &Rotor{notch: 'V', position: charToIndex('U')}
	middle := &Rotor{notch: 'E', position: charToIndex('E')} // 自身のノッチにいる
	slow := &Rotor{position: 0}
	e := &Enigma{rotors: []*Rotor{fast, middle, slow}}

	e.stepRotors()

	if fast.position != charToIndex('V') {
		t.Errorf("fast rotor should always step; got position %q, want %q", indexToChar(fast.position), 'V')
	}
}
