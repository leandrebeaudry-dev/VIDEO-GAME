package combatsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// OpenCombatInventory gère l'inventaire et l'utilisation des consommables/armes en combat
func OpenCombatInventory(p *playersystem.Player) bool {
	if len(p.Inventory) == 0 {
		fmt.Println("\nVotre inventaire est vide !")
		return false
	}

	fmt.Println("\n--- INVENTAIRE DE COMBAT ---")
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	fmt.Print("Quel objet utiliser ? : ")

	var choice int
	fmt.Scanln(&choice)

	if choice <= 0 || choice > len(p.Inventory) {
		return false
	}

	index := choice - 1
	selectedItem := p.Inventory[index]

	switch selectedItem {
	case "Potion de soin":
		if p.CurrentHP >= p.MaxHP {
			fmt.Println("\nVos PV sont déjà au maximum !")
			return false
		}
		p.CurrentHP += 40
		if p.CurrentHP > p.MaxHP {
			p.CurrentHP = p.MaxHP
		}
		p.RemoveItemFromInventory(index)
		fmt.Printf("\nVous buvez une Potion de soin ! PV : %d / %d\n", p.CurrentHP, p.MaxHP)
		return true

	case "Hache de Kratos":
		p.Atk += 25
		p.RemoveItemFromInventory(index)
		fmt.Printf("\nVous brandissez la Hache de Kratos ! Votre attaque augmente de +25 (Total : %d) !\n", p.Atk)
		return true

	case "Épée en fer":
		p.Atk += 10
		p.RemoveItemFromInventory(index)
		fmt.Printf("\nVous équipez l'Épée en fer ! Votre attaque augmente de +10 (Total : %d) !\n", p.Atk)
		return true

	default:
		fmt.Printf("\nL'objet %s ne peut pas être utilisé en combat.\n", selectedItem)
		return false
	}
}
