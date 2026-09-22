package blacksmithsystem

import (
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func StratBlacksmithSession(p *playersystem.Player) {
	fmt.Println("\n==========================================")
	fmt.Println("       LE COMPTOIR D'ACHIM LE Forgeron  ")
	fmt.Println("==========================================")

	dialogue := "Achim : « Halt, donne ta bourse et prends ton dû !\n        Ici, tout est taillé dans le fer et trempé dans le feu »\n"
	for _, char := range dialogue {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}
	for {
		fmt.Printf("\n--- Vos Pièces d'Or : %d Po ---\n", p.Gold)
		fmt.Println("1. Hache de kratos       (80 Po, +17 ATK)")
		fmt.Println("2. Épée d'Eden       (50 Po, +10 ATK)")
		fmt.Println("3. Marteau de Thor      (60 Po, +12 ATK)")
		fmt.Println("4. Lance d'Achille     (30 Po, +5 ATK)")
		fmt.Println("5. Quitter la boutique")
		fmt.Print("Achim : « Que souhaites-tu acheter ? » : ")

	var choice int
		fmt.Scanln(&choice)

		var price int

		switch choice {
		case 1:
			if price >= 80 {
			p.Gold >= price 
			p.Gold -= price 
			p.Inventory = append(p.Inventory, "Hache de Kratos")
			Atk = +17
			fmt.Println("Achim : choix judicieux, mais attention, un grand pouvoir implique de grande responsabilité")
		} else {
			fmt.Println("Achim : escroc, tu n'as point bourse à me donner")
		}

		case 2:
			if price >= 50 {
				p.Gold >= price 
				p.Gold -= price
				p.Inventory = append(p.Inventory, "Epée d'Eden")
				Atk = +10
				fmt.Println("Achim : Une arme pleine de sagesse mais destructrice")
		} else {
			fmt.Println("Achim : escroc, tu n'as point bourse à me donner")
		}
		
		case 3:
	}
}

