# AGENTS.md

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
