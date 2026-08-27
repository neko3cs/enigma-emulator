package enigma

// EnigmaKey は1台のエニグマ機を組み立てるために必要な設定一式を保持する。
// 使用するローターの種類と順序、初期位置、リングセッティング、
// プラグボードの配線、リフレクターの種類を含む。
type EnigmaKey struct {
	RotorOrder     []string
	RotorPositions string
	RingSettings   []int
	PlugboardPairs []string
	ReflectorType  string
}
