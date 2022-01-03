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
      s.Content = remove(s.Content, p)
      break
    }
    if s.Content[p].PositionX <= -10{
      s.Content = remove(s.Content, p)
      break
    }
    if s.Content[p].PositionY >= float64(config.General.WindowSizeY){
      s.Content = remove(s.Content, p)
      break
    }
    if s.Content[p].PositionY <= -10{
      s.Content[p].PositionY = float64(config.General.WindowSizeY)-2
    }
  }
  log.Println(len(s.Content))

  s.Spawnrate += config.General.SpawnRate
  for s.Spawnrate >= 1{
    s.Spawnrate -= 1
    rand.Seed(time.Now().UnixNano())
    spdX := rand.Float64()
    spdX -= 0.5
    s.Content = append(s.Content, Particle{
      PositionX: float64(config.General.SpawnX),
      PositionY: float64(config.General.SpawnY),
      ScaleX:	1, ScaleY: 1,
      ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
      Opacity: 1,
      SpeedX: spdX * 10, SpeedY: 5,
    })
  }
}



func remove(t []Particle, s int) (l []Particle){
  l = t[:s]
  l = append(l, t[s+1:]...)
  return l
}



















//
