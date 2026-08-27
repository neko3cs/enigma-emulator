package enigma

import "testing"

func TestRotorForwardBackwardAreInverses(t *testing.T) {
	wiring, notch := getRotorSpec("I")

	for position := 0; position < alphabetSize; position++ {
		for ring := 0; ring < alphabetSize; ring++ {
			r := &Rotor{wiring: wiring, notch: notch, position: position, ringSetting: ring}
			for c := byte('A'); c <= 'Z'; c++ {
				got := r.Backward(r.Forward(c))
				if got != c {
					t.Fatalf("position=%d ring=%d: Backward(Forward(%q)) = %q, want %q", position, ring, c, got, c)
				}
			}
		}
	}
}

func TestRotorAtNotch(t *testing.T) {
	r := &Rotor{notch: 'V', position: charToIndex('V')}
	if !r.AtNotch() {
		t.Errorf("expected rotor at position V to be at notch V")
	}

	r.position = charToIndex('U')
	if r.AtNotch() {
		t.Errorf("expected rotor at position U not to be at notch V")
	}
}

func TestRotorStepWrapsAround(t *testing.T) {
	r := &Rotor{position: alphabetSize - 1}
	r.Step()
	if r.position != 0 {
		t.Errorf("Step() from Z should wrap to A (0), got position %d", r.position)
	}
}
