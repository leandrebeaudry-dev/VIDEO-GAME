package blacksmithsystem

import (
	"fmt"
	"time"
)

func SpeakText(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

func WelcomeDialogue() {
	fmt.Println("\n==============================")
	fmt.Println("BIENVENUE CHEZ HAKIM LE FORGERON")
	fmt.Println("==============================")
	SpeakText("Hakim : « Bienvenue chez le forgeron, voyageur ! »")
	SpeakText("Hakim : « Que puis-je afillé pour vous ? »\n")
}
