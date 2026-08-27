package enigma

import "testing"

// TestStepRotorsDoubleStepping は、標準的なエニグマのダブルステッピング
// 挙動を固定シーケンスで回帰的に確認する。
//
// 観点: rotors III(fast)/II(middle)/I(slow)、位置AAAから100回ステップ
// させた結果のローター位置が、既知の正しいアルゴリズムでシミュレートした
// 結果（WEA）と一致すること。
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

// TestFastRotorAlwaysSteps は、右ローターの回転が省略されないことを
// 確認する回帰テスト。
//
// 観点: 中央ローターが自身のノッチにある打鍵（ダブルステップ）でも、
// 右（高速）ローターが必ず回転すること。過去にここが抜けていたバグの
// 再発防止（AGENTS.md の Incidents 参照）。
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
