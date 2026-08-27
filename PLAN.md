# PLAN.md

## Next
- Verify `rotors[0]` (rightmost/fast rotor) stepping behavior in `machine.go:EncryptChar` against known Enigma test vectors, then file a GitHub Issue if confirmed as a bug.

## Undecided
- Whether `rotors[0]` should step unconditionally on every keystroke (real Enigma behavior), independent of the `rotors[1]` notch branch. Currently `rotors[0].Step()` only runs in the `else` branch of the notch check in `machine.go`, so a keystroke that triggers the `rotors[1]`-notch branch skips advancing `rotors[0]`. Not yet confirmed as a bug — file a GitHub Issue once verified.
