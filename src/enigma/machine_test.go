package enigma

import "testing"

// testKey は、特定の鍵設定にこだわらないテストで共通して使う
// 固定の EnigmaKey を返す。
func testKey() EnigmaKey {
	return EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "ABC",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: []string{"AQ", "EP", "TY"},
		ReflectorType:  "B",
	}
}

// TestEncryptDecryptRoundTrip はエニグマ暗号の核となる性質を確認する。
//
// 観点: 同じ鍵設定の2台の機械で暗号化→復号を行うと元の文字列に一致すること
// （対称性）。また、暗号文が平文と異なる文字列になっていること。
func TestEncryptDecryptRoundTrip(t *testing.T) {
	plain := "HELLOWORLDTHISISALONGERTESTSTRINGTOCROSSNOTCHES"

	encryptor := NewEnigmaFromKey(testKey())
	cipher := encryptor.Encrypt(plain)

	decryptor := NewEnigmaFromKey(testKey())
	decrypted := decryptor.Encrypt(cipher)

	if decrypted != plain {
		t.Errorf("round trip mismatch: got %q, want %q", decrypted, plain)
	}
	if cipher == plain {
		t.Errorf("cipher text should differ from plain text, got %q", cipher)
	}
}

// TestKnownTestVector は、内部での整合性だけでなく外部の参照値との
// 一致を確認する。
//
// 観点: rotors III/II/I, ring 1/1/1, position AAA, reflector B,
// プラグボードなしで "AAAAA" を暗号化すると "BDZGO" になるという、
// 広く知られる参照テストベクターと一致すること。配線・ステッピングを
// 含めた全体の正しさを、内部実装に依存しない形で検証する。
func TestKnownTestVector(t *testing.T) {
	key := EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "AAA",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: nil,
		ReflectorType:  "B",
	}
	e := NewEnigmaFromKey(key)

	got := e.Encrypt("AAAAA")
	want := "BDZGO"
	if got != want {
		t.Errorf("Encrypt(AAAAA) = %q, want %q", got, want)
	}
}

// TestEncryptIgnoresNonAlphabet は Encrypt の入力フィルタリング挙動を
// 確認する。
//
// 観点: 記号・数字・小文字・空白など英大文字以外の文字が、出力に
// そのまま混入したり変換対象になったりしないこと。
func TestEncryptIgnoresNonAlphabet(t *testing.T) {
	e := NewEnigmaFromKey(testKey())

	got := e.Encrypt("Hello, World! 123")
	for _, c := range got {
		if c < 'A' || c > 'Z' {
			t.Fatalf("Encrypt result contains non-alphabet character: %q in %q", c, got)
		}
	}
}
