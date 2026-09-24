package combatsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// SelectAndCastSpell permet au joueur d'utiliser de la magie
func SelectAndCastSpell(p *playersystem.Player, m *Monster) bool {
	if len(p.Skill) == 0 {
		fmt.Println("\nVous ne connaissez aucun sort !")
		return false
	}

	fmt.Println("\n--- VOS SORTS ---")
	for i, spell := range p.Skill {
		fmt.Printf("%d. %s\n", i+1, spell)
	}
	fmt.Println("0. Retour")
	fmt.Print("Quel sort voulez-vous lancer ? : ")

	var choice int
	fmt.Scanln(&choice)

	if choice <= 0 || choice > len(p.Skill) {
		return false
	}

	selectedSpell := p.Skill[choice-1]

	switch selectedSpell {
	case "Coup de poing":
		cost := 10
		if p.CurrentSP < cost {
			fmt.Printf("\nPas assez de PC (%d/%d) pour Coup de poing !\n", p.CurrentSP, cost)
			return false
		}
		p.CurrentSP -= cost
		dmg := p.Atk
		m.CurrentHP -= dmg
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}
		fmt.Printf("\nVous lancez %s (-%d PC) et infligez %d dégâts au %s !\n", selectedSpell, cost, dmg, m.Name)
		return true

	case "Boule de feu":
		cost := 25
		if p.CurrentSP < cost {
			fmt.Printf("\nPas assez de PC (%d/%d) pour Boule de feu !\n", p.CurrentSP, cost)
			return false
		}
		p.CurrentSP -= cost
		dmg := p.Atk + 20
		m.CurrentHP -= dmg
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}
		fmt.Printf("\nVous lancez %s (-%d PC) et infligez %d dégâts magiques au %s !\n", selectedSpell, cost, dmg, m.Name)
		return true

	default:
		fmt.Printf("\nLe sort %s n'a pas d'effet connu.\n", selectedSpell)
		return false
	}
}
