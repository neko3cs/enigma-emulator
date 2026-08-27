package main

import (
	"bufio"
	"enigma/enigma"
	"fmt"
	"os"
	"strings"
)

// main は設定ファイル（configPath）から鍵を読み込み、標準入力から1行読み取って
// エニグマ変換（対称性により暗号化にも復号にも使える）した結果を表示する。
func main() {
	key, err := loadKey(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		return
	}

	fmt.Print("文字列を入力してください(英字のみ): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		return
	}

	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)

	enigmaMachine := enigma.NewEnigmaFromKey(key)
	cipher := enigmaMachine.Encrypt(input)

	fmt.Println("結果:", cipher)
}
