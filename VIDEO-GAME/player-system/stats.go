package playersystem

import "fmt"

// DisplayInfo affiche l'ensemble des caractéristiques du personnage
func (p *Player) DisplayInfo() {
	fmt.Println("\n--- FICHE DE PERSONNAGE ---")
	fmt.Printf("Nom       : %s\n", p.Name)
	fmt.Printf("Classe    : %s\n", p.Class)
	fmt.Printf("Niveau    : %d\n", p.Level)
	fmt.Printf("PV        : %d / %d\n", p.CurrentHP, p.MaxHP)
	fmt.Printf("Attaque   : %d\n", p.Atk)
	fmt.Printf("Or        : %d Po\n", p.Gold)
	fmt.Printf("Inventaire: %v\n", p.Inventory)
	fmt.Println("---------------------------")
}
