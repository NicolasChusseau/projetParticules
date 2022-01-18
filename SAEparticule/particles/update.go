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
      if s.Content[p].PositionX > 0 && s.Content[p].PositionX < 800 {
        for i := 0; i < s.Separation; i++ {
          s.Content[int(s.Content[p].PositionX/s.Content[0].ScaleX*18)+i-1].Vague = 10
          s.Content[int(s.Content[p].PositionX/s.Content[0].ScaleX*18)+i+17].Vague = 5
          s.Content[int(s.Content[p].PositionX/s.Content[0].ScaleX*18)+i-1].Caillou = p
          s.Content[int(s.Content[p].PositionX/s.Content[0].ScaleX*18)+i+17].Caillou = p
        }
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
				PositionX: float64(rand.Intn(98)*8),
				PositionY: -1,
				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
				Opacity: 1,
				SpeedX: 0, SpeedY: 5,
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
  				PositionX: float64(rand.Intn(100)*8),
  				PositionY: -1,
  				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
  				Opacity: 1,
  				SpeedX: 0, SpeedY: 5,
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
8:arret vague droite
9:en attente du début de la vague droite
10:en attente du début de la vague gauche*/
func Vague(s *System, p int){
  if s.Content[p].ColorBlue != 1{
    return
  }

  switch s.Content[p].Vague {
  case 10:
    s.Content[p].Vague = 1
    return
  case 9:
    s.Content[p].Vague = 5
    return
  case 8:
    s.Content[p].Vague = 0
    return
  case 7:
    s.Content[p].Vague = 8
    if s.Content[p].Caillou != 0 && proche(s.Content[p].PositionX, s.Content[s.Content[p].Caillou].PositionX)  {
      s.Content[p].SpeedY = 4*s.Content[p].ScaleY
    }else {
      s.Content[p].SpeedY = s.Content[p].ScaleY
    }
    return
  case 6:
    s.Content[p].Vague = 7
    return
  case 5:
    s.Content[p].Vague = 6
    deplacementVague(s,false,p, s.Content[p].Caillou)
    if s.Content[p].Caillou != 0 && proche(s.Content[p].PositionX, s.Content[s.Content[p].Caillou].PositionX)  {
      s.Content[p].SpeedY = -4*s.Content[p].ScaleY
    }else {
      s.Content[p].SpeedY = -s.Content[p].ScaleY
    }
    return
  case 4:
  s.Content[p].Vague = 0
  return
  case 3:
  s.Content[p].Vague = 4
  if s.Content[p].Caillou != 0 && proche(s.Content[p].PositionX, s.Content[s.Content[p].Caillou].PositionX)  {
    s.Content[p].SpeedY = 4*s.Content[p].ScaleY
  }else {
    s.Content[p].SpeedY = s.Content[p].ScaleY
  }
  return
  case 2:
    s.Content[p].Vague = 3
    return
  case 1:
    s.Content[p].Vague = 2
    deplacementVague(s,true, p, s.Content[p].Caillou)
    if s.Content[p].Caillou != 0 && proche(s.Content[p].PositionX, s.Content[s.Content[p].Caillou].PositionX)  {
      s.Content[p].SpeedY = -4*s.Content[p].ScaleY
    }else {
      s.Content[p].SpeedY = -s.Content[p].ScaleY
    }
    return
  case 0:
    s.Content[p].SpeedY = 0
    s.Content[p].Caillou = 0
    return
  }
}


func deplacementVague(s *System, gauche bool, p int, c int)  {
  if gauche && p-36 >= 0{
    s.Content[p-36].Vague = 10
    s.Content[p].Caillou = c
    return
  }
  if !gauche && p+18 < len(s.Content){
    s.Content[p+18].Vague = 9
    s.Content[p].Caillou = c
    return
  }
}

func proche(pos1, pos2 float64) bool {
  if pos1 < pos2{
    return proche(pos2, pos1)
  }
  if pos1-pos2<=float64(config.General.WindowSizeX)/4{
    return true
  }
  return false
}






















//
