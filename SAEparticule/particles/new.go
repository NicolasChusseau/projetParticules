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
		posY := rand.Float64() * float64(config.General.WindowSizeY)
		rad := (rand.Float64()*100+100)*(1-posY/float64(config.General.WindowSizeY))
		col := rand.Float64()/2
		p := Particle{
			PositionX: rand.Float64()*200+float64(config.General.WindowSizeY)/2,
			PositionYinit: posY,
			PositionY: posY,
			ScaleX:    1, ScaleY: 1,
			ColorRed: col, ColorGreen: col, ColorBlue: col,
			Opacity: 1,
			SpeedX: rand.Float64()+1, SpeedY: 0,
			Radius: rad,
		}
		sus.Content = append(sus.Content, p)
	}
	return sus
}


























/**/
