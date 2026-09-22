package merchantsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func buyitem(p *playersystem.Player, itemName string, price int, atkBonus int) {
	if p.Gold >= price {
		p.Gold -= price
		p.Inventory = append(p.Inventory, itemName)
		p.Atk = +atkBonus

		fmt.Println("Achim : Merci pour ton achat aventurier ! Tu obtiens : %s. »\n", itemName)
		if atkBonus > 0 {
			fmt.Println("(Ton attaque augmente de +%d !)\n", atkBonus)
		}

	} else {
		fmt.Println("Achim : Tu n'as pas assez d'or")
	}
}
