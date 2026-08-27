package enigma

import "testing"

func testKey() EnigmaKey {
	return EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "ABC",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: []string{"AQ", "EP", "TY"},
		ReflectorType:  "B",
	}
}

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

// AAAAA を rotors III/II/I, ring 1/1/1, position AAA, reflector B,
// プラグボードなしで暗号化すると BDZGO になる。広く知られる参照テストベクター。
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

func TestEncryptIgnoresNonAlphabet(t *testing.T) {
	e := NewEnigmaFromKey(testKey())

	got := e.Encrypt("Hello, World! 123")
	for _, c := range got {
		if c < 'A' || c > 'Z' {
			t.Fatalf("Encrypt result contains non-alphabet character: %q in %q", c, got)
		}
	}
}
