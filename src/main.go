package main

import (
	"bufio"
	"enigma/enigma"
	"fmt"
	"os"
	"strings"
)

func main() {
	key := enigma.EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "ABC",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: []string{"AQ", "EP", "TY"},
		ReflectorType:  "B",
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
