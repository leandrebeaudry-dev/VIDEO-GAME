package engine

import "math/rand"

// CalculerDegats vérifie si l'attaque est critique ou non.
// - degatsDeBase : l'attaque normale du héros ou du monstre
// - chanceCritique : pourcentage de chance de faire un crit (ex: 20 pour 20%)
func CalculerDegats(degatsDeBase int, chanceCritique int) (int, bool) {
	// Tirage aléatoire entre 0 et 99
	tirage := rand.Intn(100)

	// Si le tirage est inférieur à la chance de crit (ex: tirage < 20)
	if tirage < chanceCritique {
		// Le coup critique double les dégâts (ou x1.5)
		degatsCritiques := int(float64(degatsDeBase) * 2.0)
		return degatsCritiques, true // Retourne les gros dégâts ET un booléen 'true'
	}

	// Coup normal
	return degatsDeBase, false
}
