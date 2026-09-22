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

// InitGoblin initialise un gobelin d'entraînement
func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entraînement",
		MaxHP:      40,
		CurrentHP:  40,
		Attack:     5,
		Initiative: 3,
	}
}

// GoblinPattern gère le tour d'attaque du gobelin
func GoblinPattern(g *Monster, p *playersystem.Player, turn int) {
	var dmg int

	// Un coup puissant tous les 3 tours
	if turn%3 == 0 {
		dmg = g.Attack * 2
		fmt.Printf("Le %s prépare un coup puissant !\n", g.Name)
	} else {
		dmg = g.Attack
	}

	p.CurrentHP -= dmg
	if p.CurrentHP < 0 {
		p.CurrentHP = 0
	}

	fmt.Printf("%s inflige %d dégâts à %s !\n", g.Name, dmg, p.Name)
	fmt.Printf("Vos PV : %d / %d\n", p.CurrentHP, p.MaxHP)
}

// CharTurn gère l'action du joueur
func CharTurn(p *playersystem.Player, g *Monster) {
	fmt.Println("\n----- VOTRE TOUR -----")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		dmg := p.Atk // Utilise la vraie valeur d'attaque du joueur
		g.CurrentHP -= dmg
		if g.CurrentHP < 0 {
			g.CurrentHP = 0
		}
		fmt.Printf("\n%s attaque et inflige %d dégâts au %s !\n", p.Name, dmg, g.Name)
		fmt.Printf("PV du %s : %d / %d\n", g.Name, g.CurrentHP, g.MaxHP)
	case 2:
		fmt.Printf("\nInventaire actuel : %v\n", p.Inventory)
		fmt.Println("(Utilisation des potions à venir...)")
	default:
		fmt.Println("Choix invalide, vous perdez votre tour !")
	}
}

// TrainingFight orchestre la boucle de combat au tour par tour
func TrainingFight(p *playersystem.Player) {
	gob := InitGoblin()
	turn := 1

	fmt.Println("\n==============================")
	fmt.Println("   COMBAT D'ENTRAÎNEMENT      ")
	fmt.Println("==============================")

	// L'initiative est basée sur un seuil fixe ou sur l'attaque si l'initiative n'est pas dans Player
	playerStarts := true

	for p.CurrentHP > 0 && gob.CurrentHP > 0 {
		fmt.Printf("\n===== TOUR %d =====\n", turn)

		if playerStarts {
			CharTurn(p, &gob)
			if gob.CurrentHP <= 0 {
				fmt.Printf("\nVictoire ! Vous avez vaincu le %s !\n", gob.Name)
				p.Gold += 15 // Récompense de combat
				fmt.Println("Vous gagnez 15 pièces d'or !")
				break
			}
			GoblinPattern(&gob, p, turn)
			if p.CurrentHP <= 0 {
				fmt.Println("\nVous avez été vaincu... Retour au campement.")
				break
			}
		} else {
			GoblinPattern(&gob, p, turn)
			if p.CurrentHP <= 0 {
				fmt.Println("\nVous avez été vaincu... Retour au campement.")
				break
			}
			CharTurn(p, &gob)
			if gob.CurrentHP <= 0 {
				fmt.Printf("\nVictoire ! Vous avez vaincu le %s !\n", gob.Name)
				p.Gold += 15
				fmt.Println("Vous gagnez 15 pièces d'or !")
				break
			}
		}

		turn++
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nLe combat est terminé ! Appuyez sur Entrée pour continuer...")
}
