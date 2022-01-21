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
	sus := System{Content: []Particle{}}
	for i := 0; i < config.General.InitNumParticles; i++ {
		rand.Seed(time.Now().UnixNano())
		spdX := rand.Float64()
		spdX -= 0.5
		scale := rand.Float64()*0.01
		p := Particle{
			PositionX: rand.Float64()*float64(config.General.WindowSizeX),
			PositionY: rand.Float64()*float64(config.General.WindowSizeY),
			ScaleX:	scale, ScaleY: scale,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
			SpeedX: spdX * 2, SpeedY: 2,
		}
		sus.Content = append(sus.Content, p)
	}
	return sus
}


























/**/
