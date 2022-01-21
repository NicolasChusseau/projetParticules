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
	switch config.General.Configuration { //Choix du système créé par rapport à la configuration. Si aucune configuration n'est choisi, la configuration main est choisit par défaut
	case "main":
		return system_main()
	case "neige": //Si on choisi la configuration neige, il ne faut pas oublier de changer ParticleImage
		return system_neige()
	case "plouf":
		return system_plouf()
	case "tornade":
		return system_tornade()
	}
	return system_main()
}


/*
Création du système main
*/
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
			if config.General.Vitesse { //Si la Vitesse est défini alors on met VitesseX et VitesseY sur SpeedX et SpeedY
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

//Création du système neige
func system_neige() System {
	sus := System{Content: []Particle{}}
	rand.Seed(time.Now().UnixNano())
	var p Particle
	if config.General.RandomSpawn{ //Initialisation des positions de départ des premières particules aléatoires
		for i := 0; i < config.General.InitNumParticles; i++ {
			pX := rand.Float64() * float64(config.General.WindowSizeX)
			pY := rand.Float64() * float64(config.General.WindowSizeY)
			if config.General.Vitesse { //Si la Vitesse est défini alors on met VitesseX et VitesseY sur SpeedX et SpeedY
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

//Création du système plouf
func system_plouf() System {
	sus := System{Content: []Particle{}}
	rand.Seed(time.Now().UnixNano())
	var scalex float64 = float64(config.General.WindowSizeX)/100
	var scaley float64 = scalex
	var posX float64 = 0
	var posY float64 = float64(config.General.WindowSizeY)-scaley
	var p Particle
	for posX < float64(config.General.WindowSizeX){ //Création des particules formant l'eau. Ces particules sont placées sur toute la largeur de l'écran et sur la quart de sa hauteur
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
	sus.Separation = len(sus.Content) /100 //Cette variable représente la séparation entre chaque colonne de particules d'eau
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

//Création du système tornade
func system_tornade() System {
	sus := System{Content: []Particle{}}
	if config.General.Vitesse {
		for i := 0; i < config.General.InitNumParticles; i++ {
			rand.Seed(time.Now().UnixNano())
			posY := (rand.Float64() * (float64(config.General.WindowSizeY)+200))-100
			rad := (rand.Float64()*100)+(1-posY/float64(config.General.WindowSizeY))*100
			col := rand.Float64()/2+0.2
			if rand.Float64() < 0.5 {
				config.General.VitesseX = -config.General.VitesseX
			}
			p := Particle{
				PositionX: rand.Float64()*rad*2+(float64(config.General.WindowSizeX)/2-rad),
				PositionYinit: posY,
				PositionY: posY,
				ScaleX:    1, ScaleY: 1,
				ColorRed: col, ColorGreen: col, ColorBlue: col,
				Opacity: 1,
				SpeedX: config.General.VitesseX, SpeedY: 0,
				Radius: rad,
			}
			sus.Content = append(sus.Content, p)
		}
	}else{
		for i := 0; i < config.General.InitNumParticles; i++ {
			rand.Seed(time.Now().UnixNano())
			posY := (rand.Float64() * (float64(config.General.WindowSizeY)+200))-100
			rad := (rand.Float64()*100)+(1-posY/float64(config.General.WindowSizeY))*100
			col := rand.Float64()/2+0.2
			spdX := rand.Float64()+2
			if rand.Float64() < 0.5 {
				spdX = -spdX
			}
			p := Particle{
				PositionX: rand.Float64()*rad*2+(float64(config.General.WindowSizeX)/2-rad),
				PositionYinit: posY,
				PositionY: posY,
				ScaleX:    1, ScaleY: 1,
				ColorRed: col, ColorGreen: col, ColorBlue: col,
				Opacity: 1,
				SpeedX: spdX, SpeedY: 0,
				Radius: rad,
			}
			sus.Content = append(sus.Content, p)
		}
	}
	return sus
}
























































/**/
