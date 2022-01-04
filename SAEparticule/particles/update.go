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
func (s *System) Update() {
  for p,_ := range s.Content {
    if rand.Float64() > 0.999{
      s.Content[p].SpeedX = -s.Content[p].SpeedX
    }

    if s.Content[p].PositionX >= float64(config.General.WindowSizeX){
      s.Content = remove(s.Content, p)
      break
    }
    if s.Content[p].PositionX <= -50{
      s.Content = remove(s.Content, p)
      break
    }
    if !(s.Content[p].PositionY + float64(config.General.ImgSizeY)*s.Content[p].ScaleY >= float64(config.General.WindowSizeY)){
      s.Content[p].PositionY += s.Content[p].SpeedY
      s.Content[p].PositionX += s.Content[p].SpeedX
      r := rand.Float64()*rand.Float64()/30000
      s.Content[p].ScaleX -= r
      s.Content[p].ScaleY -= r
    }else{
      r := rand.Float64()*rand.Float64()/30000
      s.Content[p].ScaleX -= 3*r
      s.Content[p].ScaleY -= 3*r
    }

    if s.Content[p].ScaleX <= 0{
      s.Content = remove(s.Content, p)
      break
    }
  }

  s.Spawnrate += config.General.SpawnRate
  for s.Spawnrate >= 1{
    s.Spawnrate -= 1
    rand.Seed(time.Now().UnixNano())
    spdX := rand.Float64()
    spdX -= 0.5
    scale := rand.Float64()*0.04
    s.Content = append(s.Content, Particle{
      PositionX: rand.Float64()*float64(config.General.WindowSizeX),
      PositionY: -50,
      ScaleX:	scale, ScaleY: scale,
      ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
      Opacity: 1,
      SpeedX: spdX * 2, SpeedY: 2,
    })
  }
  log.Print(len(s.Content))
}

func remove(t []Particle, s int) (l []Particle){
  if s > len(t)/2{
    l = t[s+1:]
    l = append(l, t[:s]...)
  }else{
    l = t[:s]
    l = append(l, t[s+1:]...)
  }
  return l
}

func collision(t []Particle, s int) int{
  return 0
}
