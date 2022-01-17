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

    if !(s.Content[p].PositionY + float64(config.General.ImgSizeY)*s.Content[p].ScaleY >= float64(config.General.WindowSizeY)){
      s.Content[p].PositionY += s.Content[p].SpeedY
      s.Content[p].PositionX += s.Content[p].SpeedX
    }
    r := rand.Float64()*rand.Float64()/30000
    s.Content[p].ScaleX -= r
    s.Content[p].ScaleY -= r

    if s.Content[p].ScaleX <= 0{
      s.Spawnrate -= 1
      rand.Seed(time.Now().UnixNano())
      spdX := rand.Float64()
      spdX -= 0.5
      scale := rand.Float64()*0.01
      s.Content[p] = Particle{
        PositionX: rand.Float64()*float64(config.General.WindowSizeX),
        PositionY: -50,
        ScaleX:	scale, ScaleY: scale,
        ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
        Opacity: 1,
        SpeedX: spdX * 2, SpeedY: 2,
      }
    }
  }

  s.Spawnrate += config.General.SpawnRate
  for s.Spawnrate >= 1{
    s.Spawnrate -= 1
    rand.Seed(time.Now().UnixNano())
    spdX := rand.Float64()
    spdX -= 0.5
    scale := rand.Float64()*0.01
    s.Content = append(s.Content, Particle{
      PositionX: rand.Float64()*(2*float64(config.General.WindowSizeX))-float64(config.General.WindowSizeX),
      PositionY: -50,
      ScaleX:	scale, ScaleY: scale,
      ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
      Opacity: 1,
      SpeedX: spdX * 2, SpeedY: 2,
    })
  }
  log.Print(len(s.Content))
}
