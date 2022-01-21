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
	switch config.General.Configuration {
	case "main":
		return system_main()
	case "neige": //Dans ce cas, les particules sont recyclées dès qu'elles fondent. De ce fait, le MaxParticles est obselètes. C'est pourquoi on les met à 666666666666. La taille fixé ne sera pas respecté car l'image utilisé ne fait pas 1px par 1px.
	  config.General.WindowTitle = "Il neige !"
	  config.General.ParticleImage = "assets/flocon.png"
		config.General.MaxParticles = 666666666666
		return system_neige()
	case "plouf":
		return system_plouf()
	case "tornade":
		return system_tornade()
	}
		return system_main()
}



func system_main() System {
	sus := System{Content: []Particle{}}
	rand.Seed(time.Now().UnixNano())
	spdX := rand.Float64()
	if rand.Float64() > 0.5{
		spdX = -spdX
	}
	var p Particle
	if config.General.RandomSpawn{ //Initialisation des positions de départ des premières particules aléatoires
		for i := 0; i < config.General.InitNumParticles; i++ {
			if config.General.Vitesse {
				p = Particle{
					PositionX: rand.Float64() * float64(config.General.WindowSizeX),
					PositionY: rand.Float64() * float64(config.General.WindowSizeY),
					ScaleX:    config.General.Taille, ScaleY: config.General.Taille,
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
					Vie:config.General.TempsVie,
				}
			}else {
				p = Particle{
					PositionX: rand.Float64() * float64(config.General.WindowSizeX),
					PositionY: rand.Float64() * float64(config.General.WindowSizeY),
					ScaleX:    config.General.Taille, ScaleY: config.General.Taille,
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: spdX, SpeedY: -5,
					Vie:config.General.TempsVie,
				}
			}
			sus.Content = append(sus.Content, p)
		}
	}else{
		for i := 0; i < config.General.InitNumParticles; i++ {
			if config.General.Vitesse {
				p = Particle{
					PositionX: rand.Float64() * float64(config.General.WindowSizeX),
					PositionY: rand.Float64() * float64(config.General.WindowSizeY),
					ScaleX:    config.General.Taille, ScaleY: config.General.Taille,
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
					Vie:config.General.TempsVie,
				}
			}else {
				p = Particle{
					PositionX: float64(config.General.SpawnX),
					PositionY: float64(config.General.SpawnY),
					ScaleX:	config.General.Taille, ScaleY: config.General.Taille,
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: spdX * 5, SpeedY: -5,
					Vie:config.General.TempsVie,
				}
			}
			sus.Content = append(sus.Content, p)
		}
	}
	return sus
}

func system_neige() System {
	sus := System{Content: []Particle{}}
	rand.Seed(time.Now().UnixNano())
	var p Particle
	if config.General.RandomSpawn{
		for i := 0; i < config.General.InitNumParticles; i++ {
			pX := rand.Float64() * float64(config.General.WindowSizeX)
			pY := rand.Float64() * float64(config.General.WindowSizeY)
			if config.General.Vitesse {
				p = Particle{
					PositionX: pX,
					PositionY: pY,
					ScaleX: 1-(pY*100/float64(config.General.WindowSizeY)), ScaleY: 1-(pY*100/float64(config.General.WindowSizeY)),
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
					Vie: config.General.TempsVie,
				}
			}else {
				spdX := rand.Float64()
				if rand.Float64() > 0.5{
					spdX = -spdX
				}
				p = Particle{
					PositionX: pX,
					PositionY: pY,
					ScaleX: 1-(pY*100/float64(config.General.WindowSizeY)), ScaleY: 1-(pY*100/float64(config.General.WindowSizeY)),
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: spdX, SpeedY: 1,
					Vie: config.General.TempsVie,
				}
			}
			sus.Content = append(sus.Content, p)
		}
	}else{
		for i := 0; i < config.General.InitNumParticles; i++ {
			pX := rand.Float64() * float64(config.General.WindowSizeX)
			pY := rand.Float64() * float64(config.General.WindowSizeY)
			if config.General.Vitesse {
				p = Particle{
					PositionX: pX,
					PositionY: pY,
					ScaleX: 1-(pY*100/float64(config.General.WindowSizeY)), ScaleY: 1-(pY*100/float64(config.General.WindowSizeY)),
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
					Vie: config.General.TempsVie,
				}
			}else {
				spdX := rand.Float64()
				if rand.Float64() > 0.5{
					spdX = -spdX
				}
				p = Particle{
					PositionX: float64(config.General.SpawnX),
					PositionY: float64(config.General.SpawnY),
					ScaleX:	1-(pY*100/float64(config.General.WindowSizeY)), ScaleY: 1-(pY*100/float64(config.General.WindowSizeY)),
					ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
					Opacity: 1,
					SpeedX: spdX * 5, SpeedY: 2,
					Vie: config.General.TempsVie,
				}
				sus.Content = append(sus.Content, p)
			}
		}
	}
	return sus
}


func system_plouf() System {
	sus := System{Content: []Particle{}}
	rand.Seed(time.Now().UnixNano())
	var scalex float64 = float64(config.General.WindowSizeX)/100
	var scaley float64 = scalex
	var posX float64 = 0
	var posY float64 = float64(config.General.WindowSizeY)-scaley
	var p Particle
	for posX < float64(config.General.WindowSizeX){
		p = Particle{
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
			if config.General.Vitesse{
				p = Particle{
					ScaleX:    3, ScaleY: 1,
					PositionX: float64(rand.Intn(100)*8),
					PositionY: -1,
					ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
					Opacity: 1,
					SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
				}
			}else {
				p = Particle{
					ScaleX:    3, ScaleY: 1,
					PositionX: float64(rand.Intn(100)*8),
					PositionY: -1,
					ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
					Opacity: 1,
					SpeedX: 0, SpeedY: 5,
				}
			}
			sus.Content = append(sus.Content, p)
		}
	}else{
		for i := 0; i < config.General.InitNumParticles; i++ {
			if config.General.Vitesse{
				p = Particle{
					PositionX: float64(config.General.SpawnX)/2,
					PositionY: -1,
					ScaleX:	3, ScaleY: 1,
					ColorRed: 0, ColorGreen: 0, ColorBlue: 1,
					Opacity: 1,
					SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
				}
			}else {
				p = Particle{
					PositionX: float64(config.General.SpawnX)/2,
					PositionY: -1,
					ScaleX:	3, ScaleY: 1,
					ColorRed: 0, ColorGreen: 0, ColorBlue: 1,
					Opacity: 1,
					SpeedX: 0, SpeedY: 2,
				}
			}
			sus.Content = append(sus.Content, p)
		}
	}
	return sus
}


func system_tornade() System {
	sus := System{Content: []Particle{}}
	return sus
}























































/**/
