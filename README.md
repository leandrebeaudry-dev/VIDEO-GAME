# VIDEO-GAME
# VIDEO-GAME

# Projet RED — Jeu en ligne de commande (CLI)

Projet réalisé dans le cadre de l’Ymmersion Ynov.  
L’objectif est de créer un mini jeu en ligne de commande en Go, composé de plusieurs systèmes :  
création de personnage, inventaire, marchand, forgeron, combat tour par tour, économie, et progression.

---

## Fonctionnalités principales

### Système de joueur (Player System)
- Création du personnage (charCreation)
- Affichage des informations (displayInfo)
- Gestion des points de vie
- Gestion des compétences (sorts)
- Gestion de l’expérience et du niveau
- Gestion de l’initiative
- Gestion de l’équipement (bonus PV)

### Système d’inventaire (Inventory System)
- Inventaire limité à 10 items (extensible)
- Utilisation d’objets :
  - Potion de vie
  - Potion de poison
  - Livre de sort
  - Potion de mana
- Ajout / retrait d’items
- Augmentation de l’inventaire (+10 slots, max 3 fois)

### Système de marchand (Merchant System)
- Achat d’objets :
  - Potion de vie
  - Potion de poison
  - Livre de sort : Boule de feu
  - Ressources (Fourrure, Peau, Cuir, Plume)
  - Augmentation d’inventaire
- Déduction de l’or
- Ajout automatique à l’inventaire

### Système de forgeron (Blacksmith System)
- Fabrication d’équipements :
  - Chapeau de l’aventurier
  - Tunique de l’aventurier
  - Bottes de l’aventurier
- Vérification des ressources
- Déduction de l’or
- Ajout à l’inventaire
- Équipement des objets (bonus PV)

### Système de combat (Combat System)
- Combat d’entraînement contre un Gobelin
- Pattern d’attaque du Gobelin :
  - 100% dégâts chaque tour
  - 200% dégâts tous les 3 tours
- Tour du joueur :
  - Attaque basique
  - Utilisation d’objets
  - Utilisation de sorts (Coup de poing, Boule de feu)
- Gestion du mana
- Fin du combat puis retour au menu

### Système du monde (World System)
- Déplacements
- PNJ
- Zones
(Module en cours de développement)

---

## Architecture du projet