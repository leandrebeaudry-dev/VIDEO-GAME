package enginecore

import (
	"fmt"
	"os"
	"red-project/combatsystem"
	"red-project/merchantsystem"
	"red-project/playersystem"
)

// MainMenu affiche le menu principal et traite les choix du joueur
func MainMenu(p *playersystem.Player) {
	for {
		fmt.Println("\n==================================")
		fmt.Println("        PROJET RED - MENU        ")
		fmt.Println("==================================")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Ouvrir l'inventaire")
		fmt.Println("3. Visiter le marchand")
		fmt.Println("4. Lancer un combat d'entraînement")
		fmt.Println("5. Quitter le jeu")
		fmt.Println("==================================")
		fmt.Print("Votre choix : ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide.")
			// Vide le buffer en cas d'erreur de saisie
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			p.DisplayInfo()
		case 2:
			p.DisplayInventory()
		case 3:
			// Nécessite la fonction dans merchant-system
			merchantsystem.OpenShop(p)
		case 4:
			// Crée un monstre et lance le combat
			goblin := combatsystem.NewGoblin()
			combatsystem.StartFight(p, &goblin)
		case 5:
			fmt.Println("\nMerci d'avoir joué ! À bientôt.")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}
