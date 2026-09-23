package combatsystem

import (
	"fmt"
	"math/rand"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// StartTrainingSession lance un entraînement sans risque ni récompense
func StartTrainingSession(p *playersystem.Player) {
	fmt.Println("\n=== CAMP D'ENTRAÎNEMENT ===")
	fmt.Println("Un Guerrier Simérien d'entraînement s'approche !")
	fmt.Println("(Aucun risque : tes PV et ton inventaire ne seront pas modifiés à la fin)")

	// Sauvegarde des PV initiaux pour tout restaurer à la fin
	savedHP := p.CurrentHP

	// Stats temporaires pour le combat d'entraînement
	dummyHP := 60
	dummyMaxHP := 60
	dummyAttackMin := 5
	dummyAttackMax := 12

	var action int

	for dummyHP > 0 && p.CurrentHP > 0 {
		fmt.Printf("\n--- VOS PV : %d/%d | SIMÉRIEN D'ENTRAÎNEMENT PV : %d/%d ---\n", p.CurrentHP, p.MaxHP, dummyHP, dummyMaxHP)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Abandonner l'entraînement")
		fmt.Print("Choix : ")
		fmt.Scanln(&action)

		if action == 2 {
			fmt.Println("\nVous mettez fin à la session d'entraînement.")
			break
		} else if action != 1 {
			fmt.Println("Choix invalide !")
			continue
		}

		// --- TOUR DU JOUEUR ---
		playerDamage := rand.Intn(10) + 10 // Dégâts entre 10 et 19
		dummyHP -= playerDamage
		if dummyHP < 0 {
			dummyHP = 0
		}
		fmt.Printf("Vous frappez le Simérien et lui infligez %d dégâts !\n", playerDamage)

		// Vérification si le mannequin est KO
		if dummyHP == 0 {
			fmt.Println("\nBravo ! Vous avez vaincu le Guerrier Simérien d'entraînement !")
			break
		}

		// --- TOUR DU SIMÉRIEN (Attaque aléatoire) ---
		simerianDamage := rand.Intn(dummyAttackMax-dummyAttackMin+1) + dummyAttackMin
		p.CurrentHP -= simerianDamage
		if p.CurrentHP < 0 {
			p.CurrentHP = 0
		}
		fmt.Printf("Le Simérien d'entraînement contre-attaque et vous inflige %d dégâts !\n", simerianDamage)

		if p.CurrentHP == 0 {
			fmt.Println("\nLe Simérien vous a terrassé... Mais ce n'était qu'un entraînement !")
			break
		}
	}

	// Restauration des PV du joueur
	p.CurrentHP = savedHP
	fmt.Println("\nFin de l'entraînement. Vous récupérez tous vos PV !")
}
