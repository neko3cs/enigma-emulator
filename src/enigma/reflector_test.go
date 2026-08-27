package enigma

import "testing"

// リフレクターの配線は「対合（involution）」でなければならない。
// つまり2回反射すると元の文字に戻る、かつ自分自身には反射しない。
func TestReflectorIsInvolution(t *testing.T) {
	for _, name := range []string{"B", "C"} {
		r := getReflector(name)
		for c := byte('A'); c <= 'Z'; c++ {
			reflected := r.Reflect(c)
			if reflected == c {
				t.Errorf("reflector %s: %q reflects to itself", name, c)
			}
			if back := r.Reflect(reflected); back != c {
				t.Errorf("reflector %s: Reflect(Reflect(%q)) = %q, want %q", name, c, back, c)
			}
		}
	}
}

func TestGetReflectorInvalidNamePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic for invalid reflector name")
		}
	}()
	getReflector("X")
}
