# AGENTS.md

## Key References
- `docs/architecture.md` — always-on design policy (tech choices, layering, invariants)
- `PLAN.md` — in-progress work, if any

## Development Rules
- Direct commits to `main` are allowed in this repo (no worktree/branch/PR required — user instruction overrides the global default).

## Commands
```bash
cd src

# Build
go build ./...

# Vet
go vet ./...

# Test
go test ./...

# Run (reads an alphabet-only string interactively from stdin, prints the Enigma-transformed result)
go run .
```

## Tacit Knowledge
- Rotor notch detection (`Rotor.AtNotch()`) must compare the rotor's position directly against `notch`, never through `wiring`. The notch is the mechanical stepping mechanism; `wiring` is the electrical letter-substitution table. Conflating the two silently breaks the double-stepping behavior with no compile/runtime error.
- `Enigma.stepRotors()` must call `fast.Step()` unconditionally, once per keystroke, after checking notches on the pre-step position. The fast (rightmost) rotor steps every keystroke in a real Enigma with no exception, including on double-step keystrokes. Verified against the canonical test vector: rotors III/II/I, ring settings 1/1/1, start position `AAA`, reflector B, no plugboard, encrypting `"AAAAA"` must yield `"BDZGO"`.

## Incidents
| Date | What went wrong | Prevention |
| :--- | :--- | :--- |
| 2026-08-27 | `Rotor.Step()` and `Enigma.EncryptChar()` checked the notch via `r.wiring[r.position]` instead of the position letter itself, deviating from real Enigma stepping (fixed in commit df38a83) | Compare rotor position directly against `notch` — see Tacit Knowledge |
| 2026-08-27 | `Enigma.stepRotors()` skipped stepping the fast rotor entirely on double-step keystrokes, and checked the fast rotor's notch after stepping it (one keystroke too late) instead of before, deviating from real Enigma stepping (fixed in commit 35dab1a) | Fast rotor always steps once per keystroke; check notches before stepping — see Tacit Knowledge |
