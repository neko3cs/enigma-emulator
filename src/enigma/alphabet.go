package enigma

// alphabetSize はこの機械が扱うアルファベットの文字数（A〜Z）。
const alphabetSize = 26

// charToIndex は英大文字（'A'〜'Z'）を0始まりのアルファベット内インデックス
// （0〜25）に変換する。
func charToIndex(c byte) int {
	return int(c - 'A')
}

// indexToChar は0始まりのアルファベット内インデックス（0〜25）を
// 英大文字に変換する。
func indexToChar(i int) byte {
	return byte(i) + 'A'
}
