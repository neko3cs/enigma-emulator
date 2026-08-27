package enigma

import "testing"

// TestReflectorIsInvolution は、実機のリフレクターが満たすべき配線上の
// 性質を確認する。
//
// 観点: リフレクターの配線が「対合（involution）」になっていること。
// つまり2回反射すると元の文字に戻ること、かつ自分自身には反射しないこと。
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

// TestGetReflectorInvalidNamePanics は不正な入力に対する挙動を確認する。
//
// 観点: 未知のリフレクター名を渡した場合に panic すること。
func TestGetReflectorInvalidNamePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("expected panic for invalid reflector name")
		}
	}()
	getReflector("X")
}
