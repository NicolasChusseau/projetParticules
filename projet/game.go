package main

import "project-particles/particles"

// game est un type qui implémente l'interface Game de la bibliothèque Ebiten
// car il dispose d'une méthode Update, d'une méthode Draw et d'une méthode
// Layout. C'est un élément de ce type qui est utilisé pour mettre à jour et
// afficher un système de particules. Ce type ne devrait pas être modifié sauf
// pour les deux dernières extensions.
type game struct {
	system particles.System
}
