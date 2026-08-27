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

# Run (reads an alphabet-only string interactively from stdin, prints the Enigma-transformed result)
go run .
```

## Tacit Knowledge
- Rotor notch detection (`Rotor.AtNotch()`) must compare the rotor's position directly against `notch`, never through `wiring`. The notch is the mechanical stepping mechanism; `wiring` is the electrical letter-substitution table. Conflating the two silently breaks the double-stepping behavior with no compile/runtime error.

## Incidents
| Date | What went wrong | Prevention |
| :--- | :--- | :--- |
| 2026-08-27 | `Rotor.Step()` and `Enigma.EncryptChar()` checked the notch via `r.wiring[r.position]` instead of the position letter itself, deviating from real Enigma stepping (fixed in commit df38a83) | Compare rotor position directly against `notch` — see Tacit Knowledge |
