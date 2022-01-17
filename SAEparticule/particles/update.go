package particles

import (
  "project-particles/config"
  "math/rand"
  "time"
  "log"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
//Dans cette fonction, il a été décidé de respecter le MaxParticles avant le SpawnRate, il est conseillé d'avoir un maximum de particule au moins 100 fois supérieur au spawnrate.
func (s *System) Update() {
  for p,_ := range s.Content { //cette boucle sert à modifier les paramètres des particules et à vérifier si elles toujours visible à l'écran
    s.Content[p].PositionX += s.Content[p].SpeedX
    s.Content[p].PositionY += s.Content[p].SpeedY
    s.Content[p].SpeedY += config.General.Gravite
    Vague(s, p)
    if EstDansLo(s.Content[p]) && s.Content[p].ColorBlue != 1 && !s.Content[p].NonVisible{
      s.Content[p].NonVisible = true
      s.Content[p].Opacity = 0
      for i := 0; i < s.Separation; i++ {
        s.Content[int(s.Content[p].PositionX)+i].Vague = 1
        s.Content[int(s.Content[p].PositionX+s.Content[p].ScaleX)+i*18].Vague = 5
      }
    }
  }

  log.Println(len(s.Content))

  s.Spawnrate += config.General.SpawnRate
  rand.Seed(time.Now().UnixNano())
  if config.General.MaxParticles > len(s.Content){  //On respecte le maximum de particules avant de respecter le spawnrate. Le maximum de particules peut être dépassé mais il ne le sera jamais plus que le SpawnRate
    for s.Spawnrate >= 1{ //Cette boucle sert à ajouter les particules en fonction du spawnrate
      s.Spawnrate -= 1
      s.Content = append(s.Content, Particle{
        ScaleX:    3, ScaleY: 1,
				PositionX: rand.Float64() * float64(config.General.WindowSizeX),
				PositionY: -1,
				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
				Opacity: 1,
				SpeedX: 0, SpeedY: 2,
      })
    }
  }else { //Cette boucle est utilisé lorsque le maximum de particules à été atteint ou dépassé. Ici, on recycle les particules qui ne sont plus visible, celles qui sont sortie de l'écran
    for s.Spawnrate >= 1{
      s.Spawnrate -= 1
      indice := 0
      for indice < len(s.Content) && !s.Content[indice].NonVisible{ //On cherche une particules qui n'est plus à l'écran
        indice ++
      }
      if indice != len(s.Content){ //Cette condiction permet de vérifier si une particule à été trouvé ou non (si indice==len(s.Content) alors toutes les particules sont encore visible à l'écran)
        s.Content[indice] = Particle{
          ScaleX:    3, ScaleY: 1,
  				PositionX: rand.Float64() * float64(config.General.WindowSizeX),
  				PositionY: -1,
  				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
  				Opacity: 1,
  				SpeedX: 0, SpeedY: 2,
        }
      }
    }
  }
}

func EstDansLo(p Particle) bool {
  if p.PositionY >= 3*float64(config.General.WindowSizeY)/4{
    return true
  }
  return false
}


/*
0:pas de vague
1:début vague à gauche
2:milieu vague gauche
3:fin vague gauche
4:arret vague gauche
5:début vague droite
6:milieu vague droite
7:fin vague droite
8:arret vague droite*/
func Vague(s *System, p int){
  if s.Content[p].ColorBlue != 1{
    return
  }

  switch s.Content[p].Vague {
  case 0:
    s.Content[p].SpeedY = 0

  case 1:
    s.Content[p].Vague = 2
    s.Content[p].SpeedY = -s.Content[p].ScaleY
    deplacementVague(s,true, p)

  case 2:
    s.Content[p].Vague = 3

  case 3:
    s.Content[p].Vague = 4
    s.Content[p].SpeedY = +s.Content[p].ScaleY

  case 4:
  s.Content[p].Vague = 0

  case 5:
    s.Content[p].Vague = 6
    s.Content[p].SpeedY = -s.Content[p].ScaleY
    deplacementVague(s,false,p)

  case 6:
    s.Content[p].Vague = 7

  case 7:
    s.Content[p].Vague = 8
    s.Content[p].SpeedY = +s.Content[p].ScaleY

  case 8:
    s.Content[p].Vague = 0
  }
}


func deplacementVague(s *System, gauche bool, p int)  {
  if gauche && p-18 >= 0{
    s.Content[p-18].Vague = 1
  }
  if !gauche && p+18 < len(s.Content){
    s.Content[p+18].Vague = 5
  }
}
























//
