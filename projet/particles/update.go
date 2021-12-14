package particles
import (
  "time"
  "project-particles/config"
  "math/rand"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
  for i, _ :=range s.Content{
    rand.Seed(time.Now().UnixNano())
    x := rand.Float64()
    y := rand.Float64()

    if rand.Float64() > rand.Float64(){
      x = -x
    }
    if rand.Float64() > rand.Float64(){
      y = -y
    }
    if s.Content[i].PositionX < 0{
      s.Content[i].PositionX = float64(config.General.WindowSizeX)-0.5
    }
    if s.Content[i].PositionX > float64(config.General.WindowSizeX){
      s.Content[i].PositionX = 0.5
    }

    if s.Content[i].PositionY < 0{
      s.Content[i].PositionY = float64(config.General.WindowSizeY)-0.5
    }
    if s.Content[i].PositionY > float64(config.General.WindowSizeY){
      s.Content[i].PositionY = 0.5
    }
    s.Content[i].PositionX += x*10
    s.Content[i].PositionY += y*10


  }

  cpt := config.General.SpawnRate
  for cpt >= float64(1){
    s.Content = append(s.Content,Particle{
      PositionX: rand.Float64() * float64(config.General.WindowSizeX),
      PositionY: rand.Float64() * float64(config.General.WindowSizeY),
      ScaleX:    1, ScaleY: 1,
      ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
      Opacity: 0.7,
    })
    cpt -= float64(1)
  }

  if rand.Float64() < cpt{
    s.Content = append(s.Content,Particle{
      PositionX: rand.Float64() * float64(config.General.WindowSizeX),
      PositionY: rand.Float64() * float64(config.General.WindowSizeY),
      ScaleX:    1, ScaleY: 1,
      ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
      Opacity: 0.7,
  })
}


  if rand.Float64() <0.8{
    s.Content = s.Content[:len(s.Content)-1]
  }
}






































//
