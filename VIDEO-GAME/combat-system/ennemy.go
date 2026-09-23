package combatsystem

import (
	"fmt"
	"math/rand"
	"time"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func InitSimerianWarrior() Monster {
	return Monster{
		Name:      "Guerrier Simérien",
		MaxHP:     70,
		CurrentHP: 70,
		Attack:    12,
	}
}

func SimerianRandomAttack(m *Monster, p *playersystem.Player) {
	rand.Seed(time.Now().UnixNano())
	attackType := rand.Intn(4) // Choix entre 0, 1, 2 et 3

	var dmg int
	var attackName string

	switch attackType {
	case 0:
		attackName = "Coup d'épée direct"
		dmg = m.Attack
	case 1:
		attackName = "Charge simérienne brutale"
		dmg = m.Attack + 6
	case 2:
		attackName = "Coup de pommeau désorientant"
		dmg = m.Attack - 3
	case 3:
		attackName = "Estoc sanglante critique"
		dmg = m.Attack + 10
	}

	p.CurrentHP -= dmg
	if p.CurrentHP < 0 {
		p.CurrentHP = 0
	}

	fmt.Printf("\nLe %s utilise [%s] !\n", m.Name, attackName)
	fmt.Printf("Il vous inflige %d dégâts. Vos PV : %d / %d\n", dmg, p.CurrentHP, p.MaxHP)
}
