package blacksmithsystem

import (
	"fmt"
	"time"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func CountItem(inventory []string, item string) int {
	count := 0
	for _, i := range inventory {
		if i == item {
			count++
		}
	}
	return count
}

func RemoveItem(inventory []string, item string, amount int) []string {
	for i := 0; i < amount; i++ {
		for index, v := range inventory {
			if v == item {
				inventory = append(inventory[:index], inventory[index+1:]...)
				break
			}
		}
	}
	return inventory
}

func CraftItem(p *playersystem.Player, itemName string, recipe map[string]int) {
	maxInventorySize := 10
	if len(p.Inventory) >= maxInventorySize {
		fmt.Printf("\nHakim : « Ton sac est plein ! (%d/%d objets). Fais de la place d'abord. »\n", len(p.Inventory), maxInventorySize)
		return
	}

	for mat, reqCount := range recipe {
		currentCount := CountItem(p.Inventory, mat)
		if currentCount < reqCount {
			fmt.Printf("\nHakim : « Il te manque des matériaux ! Requis : %d x %s (tu en as %d). »\n", reqCount, mat, currentCount)
			return
		}
	}

	for mat, reqCount := range recipe {
		p.Inventory = RemoveItem(p.Inventory, mat, reqCount)
	}

	p.Inventory = append(p.Inventory, itemName)
	fmt.Printf("\nHakim : « Et voilà ! Ton %s est forgé avec succès ! »\n", itemName)
}

func CraftMenu(p *playersystem.Player) {
	for {
		fmt.Println("\n--- ATELIER DE FABRICATION D'Hakim ---")
		fmt.Println("1. Chapeau de l'aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println("2. Tunique de l'aventurier (2 Fourrure de loup, 1 Peau de Troll)")
		fmt.Println("3. Bottes de l'aventurier  (1 Fourrure de loup, 1 Cuir de Sanglier)")
		fmt.Println("0. Retour au comptoir")
		fmt.Print("Hakim : « Que veux-tu que je te forge ? » : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			recipe := map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}
			CraftItem(p, "Chapeau de l'aventurier", recipe)
		case 2:
			recipe := map[string]int{"Fourrure de loup": 2, "Peau de Troll": 1}
			CraftItem(p, "Tunique de l'aventurier", recipe)
		case 3:
			recipe := map[string]int{"Fourrure de loup": 1, "Cuir de Sanglier": 1}
			CraftItem(p, "Bottes de l'aventurier", recipe)
		case 0:
			return
		default:
			fmt.Println("\nHakim : « Je ne sais pas forger ça ! »")
		}
	}
}

func StartBlacksmithSession(p *playersystem.Player) {
	fmt.Println("\n==========================================")
	fmt.Println("       LE COMPTOIR D'Hakim LE FORGERON    ")
	fmt.Println("==========================================")

	dialogue := "Hakim : « Halt, donne ta bourse et prends ton dû !\n         Ici, tout est taillé dans le fer et trempé dans le feu »\n"
	for _, char := range dialogue {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}

	for {
		fmt.Printf("\n--- Vos Pièces d'Or : %d Po ---\n", p.Gold)
		fmt.Println("1. Hache de Kratos       (80 Po, +17 ATK)")
		fmt.Println("2. Épée d'Eden           (50 Po, +10 ATK)")
		fmt.Println("3. Marteau de Thor       (60 Po, +12 ATK)")
		fmt.Println("4. Lance d'Achille       (30 Po, +5 ATK)")
		fmt.Println("5. Fabriquer un équipement (Crafting)")
		fmt.Println("6. Quitter la boutique")
		fmt.Print("Hakim : « Que souhaites-tu faire ? » : ")

		var choice int
		fmt.Scanln(&choice)

		var price int

		switch choice {
		case 1:
			price = 80
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Hache de Kratos")
				p.Atk += 17
				fmt.Println("Hakim : Choix judicieux, mais attention, un grand pouvoir implique de grandes responsabilités.")
			} else {
				fmt.Println("Hakim : Escroc, tu n'as point bourse à me donner !")
			}

		case 2:
			price = 50
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Épée d'Eden")
				p.Atk += 10
				fmt.Println("Hakim : Une arme pleine de sagesse mais destructrice.")
			} else {
				fmt.Println("Hakim : Escroc, tu n'as point bourse à me donner !")
			}

		case 3:
			price = 60
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Marteau de Thor")
				p.Atk += 12
				fmt.Println("Hakim : L'arme des Dieux du tonnerre !")
			} else {
				fmt.Println("Hakim : Escroc, tu n'as point bourse à me donner !")
			}

		case 4:
			price = 30
			if p.Gold >= price {
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Lance d'Achille")
				p.Atk += 5
				fmt.Println("Hakim : Gare à ton talon ahah !")
			} else {
				fmt.Println("Hakim : Escroc, tu n'as point bourse à me donner !")
			}

		case 5:
			CraftMenu(p)

		case 6:
			fmt.Println("Hakim : Au revoir sang-mêlé(e) !")
			return

		default:
			fmt.Println("\nHakim : « Va-t'en, je ne vends pas ça ! »")
		}
	}
}
