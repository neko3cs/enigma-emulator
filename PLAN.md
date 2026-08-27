# PLAN.md

## Next
- File a GitHub Issue for the confirmed `rotors[0]` stepping bug once the user decides to (see Undecided). No other work in progress.

## Undecided
- Confirmed as a real bug (see Approach below): whether/when to file a GitHub Issue and fix it. User has twice declined to file the issue for now ("今は立てない" / "GitHub Issueは立てないって言いましたよね").

## Approach
- Verified `rotors[0]` (fast rotor) stepping in `machine.go:stepRotors()` against the standard Enigma double-stepping algorithm by simulating both side by side (position sequence for 700 keystrokes, rotors III/II/I from AAA). Two confirmed deviations:
  1. `stepRotors()` never calls `fast.Step()` in the `middle.AtNotch()` branch, so the fast rotor doesn't advance on double-step keystrokes (real Enigma: fast rotor always steps, unconditionally, every keystroke).
  2. The fast-rotor notch check happens after stepping it, so the middle rotor advances one keystroke earlier than the standard algorithm (which checks the notch on the pre-step position).
- Not fixed yet — this was investigation only, scoped separately from the "design cleanup" refactor in commit `817d024` (which intentionally preserved this behavior).

## Rejected
- Filing the GitHub Issue immediately after confirming the bug — user corrected this; the earlier "not now" decision on filing still stands and needs fresh confirmation before creating one.
