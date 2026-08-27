package enigma

import "testing"

// TestPlugboardSwap は Swap の交換ロジックを確認する。
//
// 観点: 設定したペアが双方向に正しく交換されること、
// 設定されていない文字はそのまま返ること。
func TestPlugboardSwap(t *testing.T) {
	pb := NewPlugboard([]string{"AQ", "EP"})

	cases := map[byte]byte{
		'A': 'Q',
		'Q': 'A',
		'E': 'P',
		'P': 'E',
		'Z': 'Z', // 設定されていない文字はそのまま
	}
	for in, want := range cases {
		if got := pb.Swap(in); got != want {
			t.Errorf("Swap(%q) = %q, want %q", in, got, want)
		}
	}
}
