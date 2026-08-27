package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"enigma/enigma"
)

// TestLoadKeyReadsValidConfig は loadKey の正常系を確認する。
//
// 観点: JSON設定ファイルの各フィールドが、対応する EnigmaKey の
// フィールドへ正しくデコードされること。
func TestLoadKeyReadsValidConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	content := `{
		"rotorOrder": ["III", "II", "I"],
		"rotorPositions": "ABC",
		"ringSettings": [1, 1, 1],
		"plugboardPairs": ["AQ", "EP", "TY"],
		"reflectorType": "B"
	}`
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	got, err := loadKey(path)
	if err != nil {
		t.Fatalf("loadKey returned error: %v", err)
	}

	want := enigma.EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "ABC",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: []string{"AQ", "EP", "TY"},
		ReflectorType:  "B",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("loadKey() = %+v, want %+v", got, want)
	}
}

// TestLoadKeyMissingFile は設定ファイルが存在しない場合の挙動を確認する。
//
// 観点: ファイルが存在しない場合に panic せず error を返すこと。
func TestLoadKeyMissingFile(t *testing.T) {
	_, err := loadKey(filepath.Join(t.TempDir(), "does-not-exist.json"))
	if err == nil {
		t.Errorf("expected an error for a missing config file, got nil")
	}
}

// TestLoadKeyInvalidJSON は設定ファイルの中身が不正なJSONの場合の挙動を確認する。
//
// 観点: JSONとして解析できない内容の場合に error を返すこと。
func TestLoadKeyInvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte("{not valid json"), 0o600); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	_, err := loadKey(path)
	if err == nil {
		t.Errorf("expected an error for invalid JSON, got nil")
	}
}
