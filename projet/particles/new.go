package particles

import (
	"project-particles/config"
	"math/rand"
	"time"
)
// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	sys := System{Content: []Particle{}}
	for i:=0; i<config.General.InitNumParticles; i++{
		rand.Seed(time.Now().UnixNano())
		par := Particle{
			PositionX: rand.Float64() * float64(config.General.WindowSizeX),
			PositionY: rand.Float64() * float64(config.General.WindowSizeY),
			ScaleX:    1, ScaleY: 1,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
		}
		sys.Content = append(sys.Content, par)
	}

	return sys
}
