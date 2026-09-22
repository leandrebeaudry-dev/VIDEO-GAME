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
	}
	fmt.Println("Achim : Bénis soit votre aventure champion §")
	if atkBonus > 0 {
		fmt.Println("(Achim : Ton attaque augmente de +%d !)\n", atkBonus)

	} else {
		fmt.Println("Achim : Arh, tu ne possède point assez")
	}
}
