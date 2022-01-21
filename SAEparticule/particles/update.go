package particles

import (
  "project-particles/config"
  "math/rand"
  "math"
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
  rand.Seed(time.Now().UnixNano())
  for p,_ := range s.Content { //cette boucle sert à modifier les paramètres des particules et à vérifier si elles toujours visible à l'écran
    ecartX := s.Content[p].PositionX - float64(config.General.WindowSizeX)/2
    log.Println(float64(config.General.WindowSizeX)/2)
    log.Println(float64(config.General.WindowSizeX)/2 + s.Content[p].Radius)
    log.Println(float64(config.General.WindowSizeX)/2 - s.Content[p].Radius)
    if ecartX >= float64(config.General.WindowSizeX)/2 + s.Content[p].Radius || ecartX <= float64(config.General.WindowSizeX)/2 - s.Content[p].Radius {
      s.Content[p].SpeedX = -s.Content[p].SpeedX
    }

    log.Println(ecartX)
    ecartY := math.Pow(math.Pow(s.Content[p].Radius, 2) - math.Pow(ecartX, 2), 1/2)
    if s.Content[p].SpeedX < 0 {
      s.Content[p].PositionY = s.Content[p].PositionYinit - ecartY
    }else{
      s.Content[p].PositionY = s.Content[p].PositionYinit + ecartY
    }
    s.Content[p].PositionX += s.Content[p].SpeedX

  }

  log.Println(s.Content[0])
}


func abs(x float64) float64 {
  if x < 0{
    return -x
  }
  return x
}















//
