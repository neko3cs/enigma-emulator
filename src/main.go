package main

import (
	"enigma/enigma"
	"fmt"
)

func main() {
	key := enigma.EnigmaKey{
		RotorOrder:     []string{"III", "II", "I"},
		RotorPositions: "ABC",
		RingSettings:   []int{1, 1, 1},
		PlugboardPairs: []string{"AQ", "EP", "TY"},
		ReflectorType:  "B",
	}

	message := "HELLOWORLD"

	enigma1 := enigma.NewEnigmaFromKey(key)
	cipher := enigma1.Encrypt(message)
	fmt.Println("暗号文:", cipher)

	enigma2 := enigma.NewEnigmaFromKey(key)
	plain := enigma2.Encrypt(cipher)
	fmt.Println("復号文:", plain)
}
