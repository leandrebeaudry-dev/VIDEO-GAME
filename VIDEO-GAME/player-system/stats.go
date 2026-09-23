package playersystem

import "fmt"

// Heal restaure les PV du joueur
func (p *Player) Heal(amount int) {
	p.CurrentHP += amount
	if p.CurrentHP > p.MaxHP {
		p.CurrentHP = p.MaxHP
	}
	fmt.Printf("Vous vous soignez de %d PV. (PV actuels: %d/%d)\n", amount, p.CurrentHP, p.MaxHP)
}

// RestoreSP réinitialise les PC au maximum
func (p *Player) RestoreSP() {
	p.CurrentSP = p.SkillPoints
}
