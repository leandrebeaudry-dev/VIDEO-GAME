package combatsystem

import (
	"fmt"
	"time"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

// StartBattleSession permet de choisir le monstre à affronter
func StartBattleSession(p *playersystem.Player) {
	fmt.Println("\n==========================================")
	fmt.Println("         ZONES DE CHASSE ET COMBAT        ")
	fmt.Println("==========================================")
	fmt.Println("1. Chasser un Corbeau Noir  (Loot : Plume de Corbeau)")
	fmt.Println("2. Chasser un Sanglier      (Loot : Cuir de Sanglier)")
	fmt.Println("3. Chasser un Loup Sauvage  (Loot : Fourrure de loup)")
	fmt.Println("4. Chasser un Troll         (Loot : Peau de Troll)")
	fmt.Println("5. Affronter le Guerrier Simérien (Boss / Vrai combat)")
	fmt.Println("0. Annuler")
	fmt.Print("Choisissez votre cible : ")

	var choice int
	fmt.Scanln(&choice)

	var enemy Monster

	switch choice {
	case 1:
		enemy = Monster{Name: "Corbeau Noir", MaxHP: 30, CurrentHP: 30, Attack: 5}
	case 2:
		enemy = Monster{Name: "Sanglier Enragé", MaxHP: 50, CurrentHP: 50, Attack: 10}
	case 3:
		enemy = Monster{Name: "Loup Sauvage", MaxHP: 70, CurrentHP: 70, Attack: 12}
	case 4:
		enemy = Monster{Name: "Troll des Montagnes", MaxHP: 120, CurrentHP: 120, Attack: 18}
	case 5:
		dialogue := "\nUn Guerrier Simérien surgit de la vallée...\n"
		for _, char := range dialogue {
			fmt.Printf("%c", char)
			time.Sleep(20 * time.Millisecond)
		}
		enemy = InitSimerianWarrior()
	default:
		fmt.Println("Vous faites demi-tour.")
		return
	}

	// Lance le combat
	Battle(p, &enemy, false)
}

// OpenCombatInventory gère l'inventaire et les équipements durant le tour
func OpenCombatInventory(p *playersystem.Player) bool {
	if len(p.Inventory) == 0 {
		fmt.Println("\nVotre inventaire est vide !")
		return false
	}

	fmt.Println("\n--- INVENTAIRE DE COMBAT ---")
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Retour")
	fmt.Print("Quel objet utiliser ? : ")

	var choice int
	fmt.Scanln(&choice)

	if choice <= 0 || choice > len(p.Inventory) {
		return false
	}

	selectedItem := p.Inventory[choice-1]

	switch selectedItem {
	case "Potion de soin":
		if p.CurrentHP >= p.MaxHP {
			fmt.Println("\nVos PV sont déjà au maximum !")
			return false
		}
		p.CurrentHP += 40
		if p.CurrentHP > p.MaxHP {
			p.CurrentHP = p.MaxHP
		}
		p.Inventory = append(p.Inventory[:choice-1], p.Inventory[choice:]...)
		fmt.Printf("\nVous buvez une Potion de soin ! PV : %d / %d\n", p.CurrentHP, p.MaxHP)
		return true

	case "Hache de Kratos":
		p.Atk += 25
		fmt.Printf("\nVous brandissez la Hache de Kratos ! Votre attaque passe à %d !\n", p.Atk)
		return true

	case "Épée en fer":
		p.Atk += 10
		fmt.Printf("\nVous équipez l'Épée en fer ! Votre attaque passe à %d !\n", p.Atk)
		return true

	default:
		fmt.Printf("\nVous utilisez %s !\n", selectedItem)
		return true
	}
}

// SelectAndCastSpell permet au joueur de choisir et de lancer un sort
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
			fmt.Printf("\nPas assez de PC (%d/%d) pour lancer Coup de poing !\n", p.CurrentSP, cost)
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
			fmt.Printf("\nPas assez de PC (%d/%d) pour lancer Boule de feu !\n", p.CurrentSP, cost)
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

// CharTurn gère l'action du joueur (Retourne true si le tour est joué, false si retour)
func CharTurn(p *playersystem.Player, m *Monster) (bool, bool) {
	hasWeapon := len(p.Inventory) > 0

	if p.CurrentSP <= 0 && !hasWeapon {
		fmt.Println("\nVous n'avez plus de Points de Compétence et aucune arme !")
		fmt.Println("Incapable de lutter, vous tombez au combat...")
		p.CurrentHP = 0
		return true, false
	}

	fmt.Println("\n----- VOTRE TOUR -----")
	fmt.Printf("PV: %d/%d | PC: %d/%d\n", p.CurrentHP, p.MaxHP, p.CurrentSP, p.SkillPoints)
	fmt.Println("1. Attaque physique basique")
	fmt.Println("2. Utiliser un Sort (Magie)")
	fmt.Println("3. Ouvrir l'inventaire")
	fmt.Println("4. Fuir")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		cost := 15
		if p.CurrentSP < cost {
			fmt.Println("\nPas assez de PC ! Attaque physique d'urgence (dégâts réduits).")
			dmg := p.Atk / 2
			m.CurrentHP -= dmg
			fmt.Printf("Vous infligez %d dégâts au %s.\n", dmg, m.Name)
		} else {
			p.CurrentSP -= cost
			dmg := p.Atk
			m.CurrentHP -= dmg
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}
			fmt.Printf("\nVous attaquez (-%d PC) et infligez %d dégâts au %s !\n", cost, dmg, m.Name)
		}
		fmt.Printf("PV du %s : %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
		return true, false

	case 2:
		used := SelectAndCastSpell(p, m)
		if used {
			fmt.Printf("PV du %s : %d / %d\n", m.Name, m.CurrentHP, m.MaxHP)
		}
		return used, false

	case 3:
		used := OpenCombatInventory(p)
		return used, false

	case 4:
		fmt.Println("\nVous prenez la fuite pour sauver votre peau !")
		p.GainXP(5)
		return true, true

	default:
		fmt.Println("\nHésitation... Vous perdez votre tour !")
		return true, false
	}
}

// Battle est le moteur universel de combat
func Battle(p *playersystem.Player, enemy *Monster, isTraining bool) {
	defer func() {
		fmt.Printf("\n[ Fin du combat - PV : %d/%d | PC : %d/%d ]\n", p.CurrentHP, p.MaxHP, p.CurrentSP, p.SkillPoints)
	}()

	turn := 1

	for p.CurrentHP > 0 && enemy.CurrentHP > 0 {
		fmt.Printf("\n===== TOUR %d =====\n", turn)

		turnExecuted, fled := CharTurn(p, enemy)

		if !turnExecuted {
			continue
		}

		if fled {
			break
		}

		// --- VICTOIRE DU JOUEUR ---
		if enemy.CurrentHP <= 0 {
			fmt.Printf("\nVictoire ! Vous avez vaincu le %s !\n", enemy.Name)

			// Soin de +20 PV
			p.CurrentHP += 20
			if p.CurrentHP > p.MaxHP {
				p.CurrentHP = p.MaxHP
			}
			fmt.Printf("Vous récupérez 20 PV après votre victoire ! (PV actuels : %d/%d)\n", p.CurrentHP, p.MaxHP)

			// --- DISTRIBUTION DU MATERIAL DE CRAFT (LOOT) ---
			var itemLooted string
			switch enemy.Name {
			case "Corbeau Noir":
				itemLooted = "Plume de Corbeau"
			case "Sanglier Enragé":
				itemLooted = "Cuir de Sanglier"
			case "Loup Sauvage":
				itemLooted = "Fourrure de loup"
			case "Troll des Montagnes":
				itemLooted = "Peau de Troll"
			}

			if itemLooted != "" {
				p.Inventory = append(p.Inventory, itemLooted)
				fmt.Printf("📦 Vous ramassez sur la dépouille : %s !\n", itemLooted)
			}

			// Gain d'or et d'XP
			goldEarned := 15
			p.Gold += goldEarned
			fmt.Printf("Vous obtenez %d pièces d'or ! (Total : %d Po)\n", goldEarned, p.Gold)
			p.GainXP(25)
			break
		}

		// PAR CECI :
		if isTraining {
			GoblinPattern(enemy, p, turn) // <--- On appelle GoblinPattern avec la variable turn
		} else {
			SimerianRandomAttack(enemy, p)
		}

		// --- DÉFAITE DU JOUEUR ---
		if p.CurrentHP <= 0 {
			fmt.Println("\nVous avez été vaincu... Vos alliés vous ramènent au campement.")
			p.CurrentHP = p.MaxHP / 2
			fmt.Printf("Vous reprenez vos esprits avec 50%% de vos PV max (%d/%d PV).\n", p.CurrentHP, p.MaxHP)
			break
		}

		turn++
		time.Sleep(1 * time.Second)
	}
}
