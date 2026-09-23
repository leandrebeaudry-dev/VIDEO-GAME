package blacksmithsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func buyitem(p *playersystem.Player, itemName string, price int, atkBonus int) {
	if p.Gold >= price {
		p.Gold -= price
		p.Inventory = append(p.Inventory, itemName)
		p.Atk = +atkBonus
		fmt.Println("Hakim : Bénis soit votre aventure champion §")
	} else {
		fmt.Println("Hakim : Arh, tu ne possède point assez")
	}
	if atkBonus > 0 {
		fmt.Printf("(Hakim : Ton attaque augmente de +%d !)\n", atkBonus)

	}
}
