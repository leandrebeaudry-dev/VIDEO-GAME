package playersystem

import "fmt"

// Equipment contient les équipements actuellement portés
type Equipment struct {
	Head  string
	Body  string
	Boots string
}

type Player struct {
	Name        string
	Class       string
	Level       int
	XP          int
	MaxXP       int
	MaxHP       int
	CurrentHP   int
	Atk         int
	Gold        int
	SkillPoints int
	CurrentSP   int
	Inventory   []string
	Skill       []string
	Equipped    Equipment
}

// CharacterCreation gère la création de personnage
func CharacterCreation() *Player {
	var name string
	var classChoice int

	fmt.Print("Entrez le nom de votre héros : ")
	fmt.Scanln(&name)

	fmt.Println("\nChoisissez votre classe :")
	fmt.Println("1. Guerrier Celte (PV: 120 | ATK: 16 | PC: 40)")
	fmt.Println("2. The Rock        (PV:  90 | ATK: 20 | PC: 90)")
	fmt.Println("3. Voleur          (PV:  95 | ATK: 18 | PC:  50)")
	fmt.Println("4. Mage            (PV:  80 | ATK: 10 | PC:  90)")

	fmt.Print("Votre choix (1-4) : ")
	fmt.Scanln(&classChoice)

	className := "Guerrier Celte"
	maxHP, atk, sp := 120, 16, 40

	switch classChoice {
	case 2:
		className = "The Rock"
		maxHP, atk, sp = 90, 20, 90
	case 3:
		className = "Voleur"
		maxHP, atk, sp = 95, 18, 50
	case 4:
		className = "Mage"
		maxHP, atk, sp = 80, 10, 90
	}

	return &Player{
		Name:        name,
		Class:       className,
		Level:       1,
		XP:          0,
		MaxXP:       100,
		MaxHP:       maxHP,
		CurrentHP:   maxHP,
		Atk:         atk,
		Gold:        100,
		SkillPoints: sp,
		CurrentSP:   sp,
		Inventory:   []string{"Potion de soin"},
		Skill:       []string{"Coup de poing"},
		Equipped:    Equipment{},
	}
}

// RemoveItemFromInventory supprime un objet à un index précis
func (p *Player) RemoveItemFromInventory(index int) {
	if index >= 0 && index < len(p.Inventory) {
		p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
	}
}

// EquipItem équipe une pièce d'armure
func (p *Player) EquipItem(itemName string) {
	itemIndex := -1
	for i, item := range p.Inventory {
		if item == itemName {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		fmt.Println("\nVous ne possédez pas cet objet dans votre inventaire !")
		return
	}

	switch itemName {
	case "Chapeau de l'aventurier":
		if p.Equipped.Head != "" {
			p.unequipItem(p.Equipped.Head)
		}
		p.Equipped.Head = itemName
		p.MaxHP += 10
		p.CurrentHP += 10
		fmt.Println("\nVous équipez le Chapeau de l'aventurier (+10 PV Max) !")

	case "Tunique de l'aventurier":
		if p.Equipped.Body != "" {
			p.unequipItem(p.Equipped.Body)
		}
		p.Equipped.Body = itemName
		p.MaxHP += 25
		p.CurrentHP += 25
		fmt.Println("\nVous équipez la Tunique de l'aventurier (+25 PV Max) !")

	case "Bottes de l'aventurier":
		if p.Equipped.Boots != "" {
			p.unequipItem(p.Equipped.Boots)
		}
		p.Equipped.Boots = itemName
		p.MaxHP += 15
		p.CurrentHP += 15
		fmt.Println("\nVous équipez les Bottes de l'aventurier (+15 PV Max) !")

	default:
		fmt.Println("\nCet objet ne peut pas être équipé comme pièce d'armure.")
		return
	}

	p.RemoveItemFromInventory(itemIndex)
}

func (p *Player) unequipItem(oldItem string) {
	switch oldItem {
	case "Chapeau de l'aventurier":
		p.MaxHP -= 10
	case "Tunique de l'aventurier":
		p.MaxHP -= 25
	case "Bottes de l'aventurier":
		p.MaxHP -= 15
	}

	if p.CurrentHP > p.MaxHP {
		p.CurrentHP = p.MaxHP
	}

	p.Inventory = append(p.Inventory, oldItem)
	fmt.Printf("Vous déséquipez %s (remis dans l'inventaire).\n", oldItem)
}

func (p *Player) SpellBook() {
	for _, spell := range p.Skill {
		if spell == "Boule de feu" {
			fmt.Println("\nVous connaissez déjà le sort Boule de feu !")
			return
		}
	}
	p.Skill = append(p.Skill, "Boule de feu")
	fmt.Println("\nVous avez appris le sort : Boule de feu !")
}

func (p *Player) DisplayInfo() {
	fmt.Println("\n--- FICHE DE PERSONNAGE ---")
	fmt.Printf("Nom        : %s\n", p.Name)
	fmt.Printf("Classe     : %s\n", p.Class)
	fmt.Printf("Niveau     : %d (%d / %d XP)\n", p.Level, p.XP, p.MaxXP)
	fmt.Printf("PV         : %d / %d\n", p.CurrentHP, p.MaxHP)
	fmt.Printf("PC         : %d / %d\n", p.CurrentSP, p.SkillPoints)
	fmt.Printf("Attaque    : %d\n", p.Atk)
	fmt.Printf("Or         : %d Po\n", p.Gold)
	fmt.Printf("Sorts      : %v\n", p.Skill)
	fmt.Printf("Équipement : Tête: [%s] | Torse: [%s] | Bottes: [%s]\n",
		p.Equipped.Head, p.Equipped.Body, p.Equipped.Boots)
	fmt.Printf("Inventaire : %v\n", p.Inventory)
	fmt.Println("---------------------------")
}

// GainXP gère le gain d'XP et la montée de niveau (supporte le multi-level-up)
func (p *Player) GainXP(amount int) {
	p.XP += amount
	fmt.Printf("Vous gagnez %d XP ! (%d / %d)\n", amount, p.XP, p.MaxXP)

	for p.XP >= p.MaxXP {
		p.Level++
		p.XP -= p.MaxXP
		p.MaxXP = int(float64(p.MaxXP) * 1.5)
		p.MaxHP += 15
		p.CurrentHP = p.MaxHP
		p.SkillPoints += 20
		p.CurrentSP = p.SkillPoints
		p.Atk += 3

		fmt.Printf("\n LEVEL UP ! Vous passez niveau %d !\n", p.Level)
		fmt.Printf("PV Max +15 | Attaque +3 | PC Max +20\n")
	}
}
