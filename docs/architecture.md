# アーキテクチャ

## 全体構成図

```mermaid
flowchart LR
    stdin[標準入力] --> main[main package]
    main -->|EnigmaKey| enigma[enigma package]
    enigma -->|変換結果| main
    main --> stdout[標準出力]
```

## 技術選定と理由（ADR）

### ADR-001: Go言語での実装

- 決定: Go言語で実装する。
- 背景: Go言語の学習そのものが本プロジェクトの目的（`README.md`）。
- 却下した選択肢: なし（言語選定自体が目的のため、他言語は検討していない）。
- トレードオフ: なし。

## レイヤー構成と依存の方向

- `main` → `enigma` の一方向依存。`enigma` パッケージは標準ライブラリ（`strings`, `bufio`, `os`, `fmt`）のみに依存し、外部モジュールへの依存はゼロ（`go.mod` より）。
- `main` は `enigma` の公開API（`EnigmaKey`, `NewEnigmaFromKey`, `Encrypt`）のみを利用する。`Rotor` / `Reflector` / `Plugboard` の内部フィールドはすべて非公開で、`enigma` パッケージ外から直接触れない。

## 不変条件・境界

- ローターのノッチ判定（`Rotor.AtNotch()`）は、ローターの現在位置そのものと `notch` を比較する。`wiring`（配線テーブル）を経由して判定してはならない。ノッチは回転位置を検知する機械的な仕組みであり、`wiring` は電気信号の文字置換という別レイヤーの仕組みのため。過去にこの2つを混同し、ダブルステッピング挙動が壊れる事故が発生している（`AGENTS.md` の Incidents、および `df38a83` を参照）。

## 共通規約

TBD — 命名規則・ログ出力方針などは明示的に定めていない。

## 非機能の実現方式

TBD — テスト方針（現状ユニットテストは未整備）を含め未確定。
