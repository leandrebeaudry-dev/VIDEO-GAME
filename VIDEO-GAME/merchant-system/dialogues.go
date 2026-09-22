package merchantsystem

import (
	"fmt"
	"time"
)

// SpeakText affiche un texte de dialogue de façon animée
func SpeakText(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

// WelcomeDialogue affiche le dialogue d'accueil d'Achim
func WelcomeDialogue() {
	fmt.Println("\n==============================")
	fmt.Println("   BIENVENUE CHEZ LE MARCHAND ")
	fmt.Println("==============================")
	SpeakText("Achim : « Bienvenue chez le marchand, voyageur ! »")
	SpeakText("Achim : « Jette un œil à ce que j'ai en stock... »\n")
}
