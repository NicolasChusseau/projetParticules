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
    s.Content[p].PositionX += s.Content[p].SpeedX
    s.Content[p].PositionY += s.Content[p].SpeedY
    if s.Content[p].PositionX >= float64(config.General.WindowSizeX){
      remove(s.Content, p)
    }
    if s.Content[p].PositionX <= -10{
      remove(s.Content, p)
    }
    if s.Content[p].PositionY >= float64(config.General.WindowSizeY){
      remove(s.Content, p)
    }
    if s.Content[p].PositionY <= -10{
      s.Content[p].PositionY = float64(config.General.WindowSizeY)-2
    }
  }
  rand.Seed(time.Now().UnixNano())
  spdX := rand.NormFloat64()
  log.Print(spdX)
  spdX -= 0.5
  s.Content = append(s.Content, Particle{
    PositionX: float64(config.General.SpawnX),
    PositionY: float64(config.General.SpawnY),
    ScaleX:	1, ScaleY: 1,
    ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
    Opacity: 1,
    SpeedX: spdX * 5, SpeedY: 2,
  })
}

func remove(slice []Particle, s int) []Particle {
    return append(slice[:s], slice[s+1:]...)
}
