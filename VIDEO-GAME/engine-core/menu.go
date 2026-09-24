package enginecore

import (
	"fmt"
	"syscall"
	"time"

	blacksmithsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/blacksmith-system"
	combatsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/combat-system"
	merchantsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/merchant-system"
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
	worldsystem "github.com/leandrebeaudry-dev/VIDEO-GAME/world-system"
)

var (
	msvcrt    = syscall.NewLazyDLL("msvcrt.dll")
	procKbHit = msvcrt.NewProc("_kbhit")
	procGetCh = msvcrt.NewProc("_getch")
)

func kbhit() bool {
	ret, _, _ := procKbHit.Call()
	return ret != 0
}

func getch() byte {
	ret, _, _ := procGetCh.Call()
	return byte(ret)
}

func PlayIntro() {
	story := "Dans la bretagne Armoricaine, le domaine de Dana reignait sur les terres de leur ancêtres les Celtes, mais un jour... tout bascula lorsque les Simériens ont envahi les terres. Vous veillez sur votre femme et votre enfant; \n Un soir, Baguiar, le fils d'Hakim le forgeron est venu vous chercher...\nVotre voyage dans la vallée de Dana commence maintenant.\n\n"

	skipped := false

	for _, char := range story {
		fmt.Printf("%c", char)

		if !skipped {
			if kbhit() {
				key := getch()
				// Si touche ESPACE (' '), ENTRÉE ('\r' ou '\n')
				if key == ' ' || key == '\r' || key == '\n' {
					skipped = true
				}
			}
			time.Sleep(40 * time.Millisecond)
		}
	}
}

func pauseAttendEntree() {
	fmt.Println("\n[ Appuyez sur Entrée pour revenir au menu principal ]")
	var pause string
	fmt.Scanln(&pause)
}

func MainMenu(p *playersystem.Player) {
	var choice int

	for {
		fmt.Println("\n==============================")
		fmt.Println("       MENU PRINCIPAL        ")
		fmt.Println("==============================")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Camp d'entrainement")
		fmt.Println("4. Combat")
		fmt.Println("5. Marchand")
		fmt.Println("6. Hakim le Forgeron")
		fmt.Println("8. Exploration")

		if p.Level >= 2 {
			fmt.Println("7. 🌟 Easter Egg")
		}

		fmt.Println("9. Quitter")
		fmt.Print("Entrez votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			p.DisplayInfo()
			pauseAttendEntree()

		case 2:
			fmt.Println("\n--- INVENTAIRE ---")
			if len(p.Inventory) == 0 {
				fmt.Println("Votre inventaire est vide.")
				pauseAttendEntree()
				continue
			}

			for i, item := range p.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}

			fmt.Println("\nVoulez-vous utiliser une Potion de soin ? (1: Oui / 2: Non)")
			var action int
			fmt.Scanln(&action)

			if action == 1 {
				foundIndex := -1
				for i, item := range p.Inventory {
					if item == "Potion de soin" {
						foundIndex = i
						break
					}
				}

				if foundIndex == -1 {
					fmt.Println("Vous n'avez pas de Potion de soin !")
				} else if p.CurrentHP >= p.MaxHP {
					fmt.Println("Vos PV sont déjà au maximum !")
				} else {
					p.CurrentHP += 50
					if p.CurrentHP > p.MaxHP {
						p.CurrentHP = p.MaxHP
					}
					p.Inventory = append(p.Inventory[:foundIndex], p.Inventory[foundIndex+1:]...)
					fmt.Printf("Vous utilisez une potion ! PV actuels : %d/%d\n", p.CurrentHP, p.MaxHP)
				}
			}
			pauseAttendEntree()

		case 3:
			combatsystem.StartTrainingSession(p)
			pauseAttendEntree()

		case 4:
			combatsystem.StartBattleSession(p)
			pauseAttendEntree()

		case 5:
			merchantsystem.StartMerchantSession(p)
			pauseAttendEntree()

		case 6:
			blacksmithsystem.StartBlacksmithSession(p)
			pauseAttendEntree()

		case 7:
			if p.Level >= 2 {
				fmt.Println("\nBravo, vous avez trouvé les créateurs du jeu, Jules et Léandre !")
			} else {
				fmt.Println("Choix invalide.")
			}
			pauseAttendEntree()

		case 8:
			worldsystem.StartNpcsSession(p)
			worldsystem.StartZonesSession(p)
			worldsystem.StartMovementSession(p)
			pauseAttendEntree()
		case 9:
			fmt.Println("\nAu revoir")
			return

		default:
			fmt.Println("Choix invalide.")
			pauseAttendEntree()
		}
	}
}
