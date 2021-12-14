package particles

import (
	"project-particles/config"
	"math/rand"
	"log"
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
	spdX := rand.Float64()
	if rand.Float64() > 0.5{
		spdX = -spdX
	}
	if config.General.RandomSpawn{
		for i := 0; i < config.General.InitNumParticles; i++ {
			rand.Seed(time.Now().UnixNano())
			p := Particle{
				PositionX: rand.Float64() * float64(config.General.WindowSizeX),
				PositionY: rand.Float64() * float64(config.General.WindowSizeY),
				ScaleX:    1, ScaleY: 1,
				ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
				Opacity: 1,
				SpeedX: spdX, SpeedY: 2,
			}
			sus.Content = append(sus.Content, p)
		}
	}else{
		log.Print("nonn")
		for i := 0; i < config.General.InitNumParticles; i++ {
			p := Particle{
				PositionX: float64(config.General.SpawnX),
				PositionY: float64(config.General.SpawnY),
				ScaleX:	1, ScaleY: 1,
				ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
				Opacity: 1,
				SpeedX: spdX * 5, SpeedY: 2,
			}
			sus.Content = append(sus.Content, p)
		}
	}
	log.Print("ui")
	return sus
}