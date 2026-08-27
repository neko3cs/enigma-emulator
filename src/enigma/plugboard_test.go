package enigma

import "testing"

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
