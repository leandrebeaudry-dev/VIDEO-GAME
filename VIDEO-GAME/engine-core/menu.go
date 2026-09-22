package enginecore

import (
	"fmt"
	"time"

	blacksmithsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/blacksmith-system"
	combatsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/combat-system"
	merchantsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/merchant-system"
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// PlayIntro affiche un texte lettre par lettre
func PlayIntro() {
	story := "Dans la bretagne Armoricaine, le domaine de Dana a été dévasté par les Simériens \n Achim est venu vous chercher...\nVotre voyage dans la valée de Dana commence maintenant.\n\n"

	for _, char := range story {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

// pauseAttendEntree met en pause jusqu'à ce que le joueur appuie sur Entrée
func pauseAttendEntree() {
	fmt.Println("\n[ Appuyez sur Entrée pour revenir au menu principal ]")
	var pause string
	fmt.Scanln(&pause)
}

// MainMenu gère la boucle principale du jeu
func MainMenu(p *playersystem.Player) {
	var choice int

	for {
		fmt.Println("\n==============================")
		fmt.Println("       MENU PRINCIPAL        ")
		fmt.Println("==============================")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Camp d'entrainement")
		fmt.Println("4. Qui sont-ils")
		fmt.Println("5. Combat")
		fmt.Println("6. Marchand")
		fmt.Println("7. Achim le Forgeron")
		fmt.Println("8. Quitter")
		fmt.Print("Entrez votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			p.DisplayInfo()
			pauseAttendEntree()

		case 2:
			fmt.Println("\n--- INVENTAIRE ---")
			if len(p.Inventory) == 0 {
				fmt.Println("Votre inventaire est vide.")
				pauseAttendEntree()
				continue
			}

			for i, item := range p.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}

			fmt.Println("\nVoulez-vous utiliser une Potion de soin ? (1: Oui / 2: Non)")
			var action int
			fmt.Scanln(&action)

			if action == 1 {
				foundIndex := -1
				for i, item := range p.Inventory {
					if item == "Potion de soin" {
						foundIndex = i
						break
					}
				}

				if foundIndex == -1 {
					fmt.Println("Vous n'avez pas de Potion de soin !")
				} else if p.CurrentHP >= p.MaxHP {
					fmt.Println("Vos PV sont déjà au maximum !")
				} else {
					p.CurrentHP += 50
					if p.CurrentHP > p.MaxHP {
						p.CurrentHP = p.MaxHP
					}
					p.Inventory = append(p.Inventory[:foundIndex], p.Inventory[foundIndex+1:]...)
					fmt.Printf("Vous utilisez une potion ! PV actuels : %d/%d\n", p.CurrentHP, p.MaxHP)
				}
			}
			pauseAttendEntree()

		case 3:
			fmt.Println("\nEN construction...")
			pauseAttendEntree()

		case 4:
			fmt.Println("\nBravo, vous avez trouver les créateurs du jeux, Jules et Léandre!")
			pauseAttendEntree()

		case 5:
			combatsystem.StartBattleSession(p)
			pauseAttendEntree()

		case 6:
			merchantsystem.StartMerchantSession(p)
			pauseAttendEntree()

		case 7:
			blacksmithsystem.StratBlacksmithSession(p)
			pauseAttendEntree()

		case 8:
			fmt.Println("\nAu revoir")
			return

		default:
			fmt.Println("Choix invalide.")
			pauseAttendEntree()
		}
	}
}
