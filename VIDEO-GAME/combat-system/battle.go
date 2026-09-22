package combatsystem

import (
	"fmt"
	"time"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// Monster représente l'ennemi du combat
type Monster struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int
}

// InitSimerianWarrior crée un Guerrier Simérien (ennemi de l'histoire)
func InitSimerianWarrior() Monster {
	return Monster{
		Name:       "Guerrier Simérien",
		MaxHP:      70,
		CurrentHP:  70,
		Attack:     12,
		Initiative: 4,
	}
}

// StartBattleSession est appelée par le menu principal pour lancer un vrai combat
func StartBattleSession(p *playersystem.Player) {
	dialogue := "\nUn obscur Guerrier Simérien surgit de l'ombre de la vallée...\n"
	for _, char := range dialogue {
		fmt.Printf("%c", char)
		time.Sleep(30 * time.Millisecond)
	}

	// Lance le combat de l'histoire
	Battle(p)
}

// SimerianPattern gère le tour d'attaque du Simérien
func SimerianPattern(m *Monster, p *playersystem.Player, turn int) {
	var dmg int

	// Attaque lourde tous les 3 tours
	if turn%3 == 0 {
		dmg = m.Attack + 6
		fmt.Printf("Le %s frappe avec rage !\n", m.Name)
	} else {
		dmg = m.Attack
	}

	p.CurrentHP -= dmg
	if p.CurrentHP < 0 {
		p.CurrentHP = 0
	}

	fmt.Printf("Le %s vous inflige %d dégâts !\n", m.Name, dmg)
	fmt.Printf("Vos PV : %d / %d\n", p.CurrentHP, p.MaxHP)
}

// CharTurn gère l'action du joueur pendant le combat
func CharTurn(p *playersystem.Player, m *Monster) {
	fmt.Println("\n----- VOTRE TOUR -----")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser une Potion de soin")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		dmg := p.Atk
		m.CurrentHP -= dmg
		if m.CurrentHP < 0 {
			m.CurrentHP = 0
		}
		fmt.Printf("\n%s attaque et inflige %d dégâts au %s !\n", p.Name, dmg, m.Name)
		fmt.Printf("PV du %s : %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)

	case 2:
		foundIndex := -1
		for i, item := range p.Inventory {
			if item == "Potion de soin" {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			fmt.Println("\nVous n'avez aucune Potion de soin dans votre inventaire !")
		} else if p.CurrentHP >= p.MaxHP {
			fmt.Println("\nVos PV sont déjà au maximum !")
		} else {
			p.CurrentHP += 40
			if p.CurrentHP > p.MaxHP {
				p.CurrentHP = p.MaxHP
			}
			p.Inventory = append(p.Inventory[:foundIndex], p.Inventory[foundIndex+1:]...)
			fmt.Printf("\nVous buvez une potion ! PV restaurés. Vos PV : %d / %d\n", p.CurrentHP, p.MaxHP)
		}

	default:
		fmt.Println("\nChoix invalide, vous hésitez et perdez votre tour !")
	}
}

// Battle gère la boucle principale du vrai combat
func Battle(p *playersystem.Player) {
	enemy := InitSimerianWarrior()
	turn := 1

	fmt.Println("\n==============================")
	fmt.Println("     COMBAT : VAL LÉE DE DANA ")
	fmt.Println("==============================")

	for p.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Printf("\n===== TOUR %d =====\n", turn)

		// Le joueur attaque en premier
		CharTurn(p, &enemy)
		if enemy.CurrentHP <= 0 {
			fmt.Printf("\nVictoire glorieuse ! Vous avez terrassé le %s !\n", enemy.Name)
			p.Gold += 35
			fmt.Println("Vous récupérez 35 pièces d'or sur le cadavre du Simérien !")
			break
		}

		// Tour de l'ennemi
		SimerianPattern(&enemy, p, turn)
		if p.CurrentHP <= 0 {
			fmt.Println("\nVous vous écroulez au combat... Le domaine de Dana succombe aux Simériens.")
			break
		}

		turn++
		time.Sleep(1 * time.Second)
	}
}
