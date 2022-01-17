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
	rand.Seed(time.Now().UnixNano())
	var scalex float64 = float64(config.General.WindowSizeX)/100
	var scaley float64 = scalex
	var posX float64 = 0
	var posY float64 = float64(config.General.WindowSizeY)-scaley
	for posX < float64(config.General.WindowSizeX){
		p := Particle{
			ScaleX:    scalex, ScaleY: scaley,
			PositionX: posX,
			PositionY: posY,
			ColorRed: 0, ColorGreen: 0, ColorBlue: 1,
			Opacity: 1,
			SpeedX: 0, SpeedY: 0,
		}
		sus.Content = append(sus.Content, p)
		posY -= scaley
		if posY <= 3*float64(config.General.WindowSizeY)/4{
			posY = float64(config.General.WindowSizeY)-scaley
			posX += scalex
		}
	}
	sus.Separation = len(sus.Content) /100
	if config.General.RandomSpawn{ //Initialisation des positions de départ des premières particules aléatoires
		for i := 0; i < config.General.InitNumParticles; i++ {
			p := Particle{
				ScaleX:    3, ScaleY: 1,
				PositionX: rand.Float64() * float64(config.General.WindowSizeX),
				PositionY: -1,
				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
				Opacity: 1,
				SpeedX: 0, SpeedY: 2,
			}
			sus.Content = append(sus.Content, p)
		}
	}else{
		for i := 0; i < config.General.InitNumParticles; i++ {
			p := Particle{
				PositionX: float64(config.General.SpawnX)/2,
				PositionY: -1,
				ScaleX:	3, ScaleY: 1,
				ColorRed: 0, ColorGreen: 0, ColorBlue: 1,
				Opacity: 1,
				SpeedX: 0, SpeedY: 2,
			}
			sus.Content = append(sus.Content, p)
		}
	}
	return sus
}


























/**/
