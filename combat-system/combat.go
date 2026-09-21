package combatsystem

import (
	"fmt"
)

// TrainingFight gère la boucle d'un combat tour par tour[cite: 1, 1]
func StartFight(p *playersystem.Player, monster *Monster) {
	fmt.Printf("\n Un %s sauvage apparaît ! (PV: %d | Attaque: %d)\n", monster.Name, monster.MaxHP, monster.Attack)

	turn := 1

	for p.CurrentHP > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n--- TOUR %d ---\n", turn)
		fmt.Printf("Vos PV : %d/%d | PV %s : %d/%d\n", p.CurrentHP, p.MaxHP, monster.Name, monster.CurrentHP, monster.MaxHP)

		// 1. Tour du joueur
		fmt.Println("1. Attaquer")
		fmt.Println("2. Utiliser une potion")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			damage := 10 // Dégâts de base du joueur
			monster.CurrentHP -= damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}
			fmt.Printf("Vous attaquez le %s et lui infligez %d dégâts !\n", monster.Name, damage)
		case 2:
			p.TakePotion()
		default:
			fmt.Println("Choix invalide, vous perdez votre tour !")
		}

		// Vérification si le monstre est mort
		if monster.CurrentHP <= 0 {
			fmt.Printf("\n Victoire ! Vous avez vaincu le %s !\n", monster.Name)
			// Gain de récompenses (ex: pièces d'or)
			p.Gold += 10
			fmt.Println("Vous gagnez 10 pièces d'or.")
			return
		}

		// 2. Tour du monstre (Pattern IA)
		// Attaque puissante (200% de dégâts) tous les 3 tours
		monsterDamage := monster.Attack
		if turn%3 == 0 {
			monsterDamage = monster.Attack * 2
			fmt.Printf("⚡ Le %s lance un coup puissant !\n", monster.Name)
		}

		p.CurrentHP -= monsterDamage
		fmt.Printf("Le %s vous attaque et vous inflige %d dégâts !\n", monster.Name, monsterDamage)

		// Vérification si le joueur est mort
		if p.CurrentHP <= 0 {
			p.CurrentHP = 0
			fmt.Println("\n Vous avez été vaincu...")
			p.CheckDeath()
			return
		}

		turn++
	}
}
