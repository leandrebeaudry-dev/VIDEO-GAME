package playersystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name      string
	Class     string
	HP        int
	MaxHP     int
	Atk       int
	Gold      int
	Inventory []string
}

func CharacterCreation() *Player {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==============================")
	fmt.Println("   CRÉATION DU PERSONNAGE     ")
	fmt.Println("==============================")

	fmt.Print("Entrez le nom de votre héros : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Aventurier"
	}

	fmt.Println("\nChoisissez une classe :")
	fmt.Println("1. Guerrier Celte (120 PV | 15 ATK)")
	fmt.Println("2. Druide (80 PV  | 25 ATK)")
	fmt.Println("3. Voleur   (100 PV | 20 ATK)")
	fmt.Println("4. Dobby    (50 PV  | 40 ATK)")

	var choice int
	fmt.Print("Votre choix (1-4) : ")
	fmt.Scanln(&choice)

	var hp, atk int
	var class, starterItem string

	switch choice {
	case 1:
		class, hp, atk, starterItem = "Guerrier Celte", 120, 15, "Épée en fer"
	case 2:
		class, hp, atk, starterItem = "Druide", 80, 25, "Bâton magique"
	case 3:
		class, hp, atk, starterItem = "Voleur", 100, 20, "Dague en acier"
	case 4:
		class, hp, atk, starterItem = "Dobby", 50, 40, "Sort d'enchantement"
	default:
		class, hp, atk, starterItem = "Guerrier", 120, 15, "Épée en fer"
	}

	return &Player{
		Name:      name,
		Class:     class,
		HP:        hp,
		MaxHP:     hp,
		Atk:       atk,
		Inventory: []string{"Potion de soin", starterItem},
	}
}

func (p *Player) ShowInfo() {
	fmt.Println("\n--- INFORMATIONS ---")
	fmt.Printf("Nom     : %s\n", p.Name)
	fmt.Printf("Classe  : %s\n", p.Class)
	fmt.Printf("Santé   : %d / %d PV\n", p.HP, p.MaxHP)
	fmt.Printf("Attaque : %d\n", p.Atk)
}
