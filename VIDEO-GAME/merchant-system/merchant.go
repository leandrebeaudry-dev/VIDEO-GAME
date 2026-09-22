package merchantsystem

import (
	"fmt"
	"time"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// StartMerchantSession gère la rencontre avec le marchand Achim
func StartMerchantSession(p *playersystem.Player) {
	fmt.Println("\n==========================================")
	fmt.Println("       LE COMPTOIR DE REÜS LE MARCHAND    ")
	fmt.Println("==========================================")

	// Animation de texte pour le dialogue
	dialogue := "Reüs: « Salutations, aventurier ! Les Simériens rôdent...\n        Prends de quoi te défendre ou te soigner ! »\n"
	for _, char := range dialogue {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}

	for {
		fmt.Printf("\n--- Vos Pièces d'Or : %d Po ---\n", p.Gold)
		fmt.Println("1. Potion de soin       (20 Po)")
		fmt.Println("2. Épée en acier        (50 Po, +10 ATK)")
		fmt.Println("3. Bâton magique        (40 Po, +7 ATK)")
		fmt.Println("4. Dague en acier       (30 Po, +5 ATK)")
		fmt.Println("5. Sort d'enchantement  (50 Po)")
		fmt.Println("6. Quitter la boutique")
		fmt.Print("Reüs : « Que souhaites-tu acheter ? » : ")

		var choice int
		fmt.Scanln(&choice)

		var price int

		switch choice {
		case 1:
			price = 20
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Potion de soin")
				fmt.Println("\nReüs : « Et voilà une Potion de soin bien fraîche ! »")
			} else {
				fmt.Println("\nReüs : « Tu n'as pas assez d'or pour ça, mon ami... »")
			}

		case 2:
			price = 50
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Épée en acier")
				p.Atk += 10
				fmt.Println("\nReüs : « Une excellente lame ! Ton attaque augmente de +10. »")
			} else {
				fmt.Println("\nReüs : « Ah, tu n'as pas assez de pièces. »")
			}

		case 3:
			price = 40
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Bâton magique")
				p.Atk += 7
				fmt.Println("\nReüs : « Un bâton très puissant ! Ton attaque augmente de +7. »")
			} else {
				fmt.Println("\nReüs : « Ah, tu n'as pas assez de pièces. »")
			}

		case 4:
			price = 30
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Dague en acier")
				p.Atk += 5
				fmt.Println("\nReüs : « Une dague bien aiguisée ! Ton attaque augmente de +5. »")
			} else {
				fmt.Println("\nReüs : « Ah, tu n'as pas assez de pièces. »")
			}

		case 5:
			price = 50
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Sort d'enchantement")
				fmt.Println("\nReüs : « Un sortilège ancien... Utilise-le avec sagesse ! »")
			} else {
				fmt.Println("\nReüs : « Ah, tu n'as pas assez de pièces. »")
			}

		case 6:
			fmt.Println("\nReüs : « Merci de ta visite ! Que les grands Guerriers Celtes te protège. »")
			return

		default:
			fmt.Println("\nReüs : « Il n'y a pas ça ici, bougre sans nom... »")
		}
	}
}
