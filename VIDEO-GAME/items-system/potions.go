package itemssystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// UseManaPotion consomme une potion de mana dans l'inventaire du joueur
func UseManaPotion(p *playersystem.Player) {
	itemIndex := -1
	for i, item := range p.Inventory {
		if item == "Potion de mana" {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		fmt.Println("\nVous n'avez pas de Potion de mana dans votre inventaire !")
		return
	}

	// Suppression de l'objet de l'inventaire
	p.Inventory = append(p.Inventory[:itemIndex], p.Inventory[itemIndex+1:]...)

	// Restauration de 30 PC / Mana
	p.CurrentSP += 30
	if p.CurrentSP > p.SkillPoints {
		p.CurrentSP = p.SkillPoints
	}

	fmt.Printf("\n🧪 Vous buvez une Potion de mana (+30 Mana) ! Mana actuel : %d / %d\n", p.CurrentSP, p.SkillPoints)
}
