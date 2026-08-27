package main

import (
	"encoding/json"
	"enigma/enigma"
	"os"
)

// configPath は、鍵設定を読み込むJSONファイルのパス（作業ディレクトリからの相対パス）。
const configPath = "config.json"

// loadKey は path のJSON設定ファイルを読み込み、enigma.EnigmaKey にデコードする。
// ここで読み込んだ設定を Enigma に渡す（依存性の注入）ことで、
// enigma パッケージ自身は設定の出所（ファイルであること）を知らずに済む。
func loadKey(path string) (enigma.EnigmaKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return enigma.EnigmaKey{}, err
	}

	var key enigma.EnigmaKey
	if err := json.Unmarshal(data, &key); err != nil {
		return enigma.EnigmaKey{}, err
	}
	return key, nil
}
