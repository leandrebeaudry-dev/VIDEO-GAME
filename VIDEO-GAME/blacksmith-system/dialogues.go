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
	fmt.Println("BIENVENUE CHEZ ACHIM LE FORGERON")
	fmt.Println("==============================")
	SpeakText("Achim : « Bienvenue chez le forgeron, voyageur ! »")
	SpeakText("Achim : « Que puis-je afillé pour vous ? »\n")
}
