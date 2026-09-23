package combatsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// InitGoblin initialise le Gobelin d'entraînement
func InitGoblin() Monster {
	return Monster{
		Name:      "Gobelin d'entraînement",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

// GoblinPattern gère l'attaque du Gobelin en fonction du numéro de tour
func GoblinPattern(m *Monster, p *playersystem.Player, turn int) {
	dmg := m.Attack // 100% de l'attaque par défaut (5 dégâts)

	// Tous les 3 tours (3, 6, 9...), il inflige 200% de son attaque
	if turn%3 == 0 {
		dmg = m.Attack * 2 // 200% de l'attaque (10 dégâts)
		fmt.Println("\n⚡ Le Gobelin prépare une attaque puissante !")
	}

	// Inflige les dégâts au joueur
	p.CurrentHP -= dmg
	if p.CurrentHP < 0 {
		p.CurrentHP = 0
	}

	// Affichage demandé par la consigne
	fmt.Printf("\n%s inflige à %s %d de dégâts\n", m.Name, p.Name, dmg)
	fmt.Printf("PV de %s : %d/%d\n", p.Name, p.CurrentHP, p.MaxHP)
}
