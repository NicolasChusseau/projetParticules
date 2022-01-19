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
    if s.Content[p].SpeedY < 0{
      s.Content[p].Opacity -= 0.01
    }else {
      s.Content[p].Rotation += 0.1
      s.Content[p].Opacity += 10
    }
    s.Content[p].NonVisible = EstNonVisible(s.Content[p])
  }
  log.Println(len(s.Content))

  s.Spawnrate += config.General.SpawnRate
  rand.Seed(time.Now().UnixNano())
  if config.General.MaxParticles > len(s.Content){  //On respecte le maximum de particules avant de respecter le spawnrate. Le maximum de particules peut être dépassé mais il ne le sera jamais plus que le SpawnRate
    for s.Spawnrate >= 1{ //Cette boucle sert à ajouter les particules en fonction du spawnrate
      s.Spawnrate -= 1
      spdX := rand.Float64()
      spdX -= 0.5
      spdY := rand.Float64() * 10
      s.Content = append(s.Content, Particle{
        PositionX: float64(config.General.SpawnX),
        PositionY: float64(config.General.SpawnY),
        ScaleX:	1, ScaleY: 1,
        ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
        Opacity: 1,
        SpeedX: spdX * 10, SpeedY: -(spdY+5),
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
        spdX := rand.Float64()
        spdX -= 0.5
        spdY := rand.Float64() * 10
        s.Content[indice] = Particle{
          PositionX: float64(config.General.SpawnX),
          PositionY: float64(config.General.SpawnY),
          ScaleX:	1, ScaleY: 1,
          ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
          Opacity: 1,
          SpeedX: spdX * 10, SpeedY: -(spdY+5),
        }
      }
    }
  }
}

func EstNonVisible(p Particle) bool {
  if p.PositionX >= float64(config.General.WindowSizeX) || p.PositionX <= -10 || p.PositionY >= float64(config.General.WindowSizeY) || p.PositionY <= -10{
    return true
  }
  return false
}


















//
