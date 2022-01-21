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
//Dans cette fonction, il a été décidé de respecter le MaxParticles avant le SpawnRate, il est conseillé d'avoir un maximum de particule au moins 200 fois supérieur au spawnrate si le temps de vie est autour de 150.
//Une fois le MaxParticles atteint, plus aucune particules n'est créé car les particules sont recyclées
func (s *System) Update() {
  switch config.General.Configuration {
  case "main":
    s.update_main()
    return
  case "neige":
    s.update_neige()
    return
  case "plouf": //Attention, si plusieurs caillou tombe dans l'eauen lmême temps, cela créerd des beugs. Il faut donc veiller à ne pas que ça arrive. Pour cela, il est conseillé de mettre InitNumParticles à maximum 1 et un spawnrate au maximum à 0.007
    s.update_plouf()
    return
  case "tornade":
    s.update_tornade()
    return
  }
  s.update_main()
}


func (s *System) update_main()  {
  for p,_ := range s.Content { //cette boucle sert à modifier les paramètres des particules et à vérifier si elles toujours visible à l'écran
    if !s.Content[p].NonVisible{
      s.Content[p].PositionX += s.Content[p].SpeedX
      s.Content[p].PositionY += s.Content[p].SpeedY
      s.Content[p].SpeedY += config.General.Gravite
      s.Content[p].Vie -= 1
      if s.Content[p].SpeedY < 0{
        s.Content[p].Opacity = 1
        s.Content[p].Rotation = 0
      }else {
          s.Content[p].Rotation += 0.1
          s.Content[p].Opacity += 10
      }
      if s.Content[p].Vie < -0 {
        s.Content[p].NonVisible = true
        s.Content[p].Opacity = -5000
      }
      if !config.General.Rebonds { // si les rebonds ne sont pas activés
        if EstNonVisible(s.Content[p]){ //test de la visibilité de la particule
          s.Content[p].NonVisible = true
          s.Content[p].Opacity = -5000 //on met une opacité négative pour que la particule ne soit plus visible à l'écran sur les bords si elle avait une rotation
        }
      }else{ //si les rebonds sont activés, on test si elle touche un bord et on inverse la vitesse correspondante (les rebonds ne sont pas forcément fluides et cohérents si la particule à une rotation)
        rebond(s, p)
      }
    }
  }

  log.Println(len(s.Content))

  s.Spawnrate += config.General.SpawnRate
  rand.Seed(time.Now().UnixNano())
  if config.General.MaxParticles > len(s.Content){  //On respecte le maximum de particules avant de respecter le spawnrate. Le maximum de particules peut être dépassé mais il ne le sera jamais plus que le SpawnRate
    for s.Spawnrate >= 1{ //Cette boucle sert à ajouter les particules en fonction du spawnrate
      s.Spawnrate -= 1
      if !config.General.Vitesse{
        spdX := rand.Float64()
        spdX -= 0.5
        spdY := rand.Float64() * 10
        s.Content = append(s.Content, Particle{
          PositionX: float64(config.General.SpawnX),
          PositionY: float64(config.General.SpawnY),
          ScaleX:	config.General.Taille, ScaleY: config.General.Taille,
          ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
          Opacity: 1,
          SpeedX: spdX * 10, SpeedY: -(spdY+5),
          Vie:config.General.TempsVie,
        })
      }else {
        s.Content = append(s.Content, Particle{
          PositionX: float64(config.General.SpawnX),
          PositionY: float64(config.General.SpawnY),
          ScaleX:	config.General.Taille, ScaleY: config.General.Taille,
          ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
          Opacity: 1,
          SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
          Vie:config.General.TempsVie,
        })
      }
    }
  }else { //Cette boucle est utilisé lorsque le maximum de particules à été atteint ou dépassé. Ici, on recycle les particules qui ne sont plus visible, celles qui sont sortie de l'écran
    for s.Spawnrate >= 1{
      s.Spawnrate -= 1
      indice := 0
      for indice < len(s.Content) && !s.Content[indice].NonVisible{ //On cherche une particules qui n'est plus à l'écran
        indice ++
      }
      if indice != len(s.Content){ //Cette condiction permet de vérifier si une particule à été trouvé ou non (si indice==len(s.Content) alors toutes les particules sont encore visible à l'écran)
        if !config.General.Vitesse{
          spdX := rand.Float64()
          spdX -= 0.5
          spdY := rand.Float64() * 10
          s.Content[indice] = Particle{
            PositionX: float64(config.General.SpawnX),
            PositionY: float64(config.General.SpawnY),
            ScaleX:	config.General.Taille, ScaleY: config.General.Taille,
            ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
            Opacity: 1,
            SpeedX: spdX * 10, SpeedY: -(spdY+5),
            Vie:config.General.TempsVie,
          }
        }else {
          s.Content[indice] = Particle{
            PositionX: float64(config.General.SpawnX),
            PositionY: float64(config.General.SpawnY),
            ScaleX:	config.General.Taille, ScaleY: config.General.Taille,
            ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
            Opacity: 1,
            SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
            Vie:config.General.TempsVie,
          }
        }
      }
    }
  }
}


func (s *System) update_neige() { //Dans cette fonction, le temps de vie d'une particules est en fonction de sa taille à l'écran donc on n'a pas besoin de décrémenter son TempsVie.
  rand.Seed(time.Now().UnixNano())
  var anti_beug_rebond float64
  for p,_ := range s.Content {
    s.Content[p].SpeedY += config.General.Gravite
    if rand.Float64() > 0.999{
      s.Content[p].SpeedX = -s.Content[p].SpeedX
    }

    if !config.General.Rebonds {
      if !(s.Content[p].PositionY + float64(config.General.ImgSizeY)*s.Content[p].ScaleY >= float64(config.General.WindowSizeY)){
        s.Content[p].PositionY += s.Content[p].SpeedY
        s.Content[p].PositionX += s.Content[p].SpeedX
      }
    }else{
      s.Content[p].PositionY += s.Content[p].SpeedY
      s.Content[p].PositionX += s.Content[p].SpeedX
      rebond(s, p)
      anti_beug_rebond = 3
    }
    r := rand.Float64()*rand.Float64()/30000
    s.Content[p].ScaleX -= r
    s.Content[p].ScaleY -= r

    if s.Content[p].ScaleX <= 0{
      s.Spawnrate -= 1
      scale := rand.Float64()*0.01
      if !config.General.Vitesse {
        spdX := rand.Float64()
        spdX -= 0.5
        s.Content[p] = Particle{
          PositionX: rand.Float64()*float64(config.General.WindowSizeX),
          PositionY: -2+anti_beug_rebond,
          ScaleX:	scale, ScaleY: scale,
          ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
          Opacity: 1,
          SpeedX: spdX * 2, SpeedY: 2,
        }
      }else {
        s.Spawnrate -= 1
        s.Content[p] = Particle{
          PositionX: rand.Float64()*float64(config.General.WindowSizeX),
          PositionY: -2+anti_beug_rebond,
          ScaleX:	scale, ScaleY: scale,
          ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
          Opacity: 1,
          SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
        }
      }
    }
  }

  s.Spawnrate += config.General.SpawnRate
  for s.Spawnrate >= 1{
    s.Spawnrate -= 1
    rand.Seed(time.Now().UnixNano())
    scale := rand.Float64()*0.01
    if !config.General.Vitesse {
      spdX := rand.Float64()
      spdX -= 0.5
      s.Content = append(s.Content, Particle{
        PositionX: rand.Float64()*(2*float64(config.General.WindowSizeX))-float64(config.General.WindowSizeX),
        PositionY: -2,
        ScaleX:	scale, ScaleY: scale,
        ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
        Opacity: 1,
        SpeedX: spdX * 2, SpeedY: 2,
      })
    }else {
      s.Content = append(s.Content, Particle{
        PositionX: rand.Float64()*(2*float64(config.General.WindowSizeX))-float64(config.General.WindowSizeX),
        PositionY: -2,
        ScaleX:	scale, ScaleY: scale,
        ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
        Opacity: 1,
        SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
      })
    }
  }
  log.Println(len(s.Content))
}

func (s *System) update_plouf() {
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
      if config.General.Vitesse{
        s.Content = append(s.Content, Particle{
          ScaleX:    3, ScaleY: 1,
  				PositionX: float64(rand.Intn(98)*8),
  				PositionY: -1,
  				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
  				Opacity: 1,
  				SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
        })
      }else{
        s.Content = append(s.Content, Particle{
          ScaleX:    3, ScaleY: 1,
  				PositionX: float64(rand.Intn(98)*8),
  				PositionY: -1,
  				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
  				Opacity: 1,
  				SpeedX: 0, SpeedY: 5,
        })
      }
    }
  }else { //Cette boucle est utilisé lorsque le maximum de particules à été atteint ou dépassé. Ici, on recycle les particules qui ne sont plus visible, celles qui sont sortie de l'écran
    for s.Spawnrate >= 1{
      s.Spawnrate -= 1
      indice := 0
      for indice < len(s.Content) && !s.Content[indice].NonVisible{ //On cherche une particules qui n'est plus à l'écran
        indice ++
      }
      if indice != len(s.Content){ //Cette condiction permet de vérifier si une particule à été trouvé ou non (si indice==len(s.Content) alors toutes les particules sont encore visible à l'écran)
        if config.General.Vitesse{
          s.Content[indice] = Particle{
            ScaleX:    3, ScaleY: 1,
    				PositionX: float64(rand.Intn(100)*8),
    				PositionY: -1,
    				ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
    				Opacity: 1,
    				SpeedX: config.General.VitesseX, SpeedY: config.General.VitesseY,
          }
        }else {
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
}


func (s *System) update_tornade() {
}







/*
Renvoie vrai si une particules est dans l'eau
*/
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

/*
Effectue les rebonds
*/
func rebond(s *System, p int)  {
  if s.Content[p].PositionX >= float64(config.General.WindowSizeX)-s.Content[p].ScaleX || s.Content[p].PositionX <= 0{
    s.Content[p].SpeedX = -s.Content[p].SpeedX
  }
  if s.Content[p].PositionY >= float64(config.General.WindowSizeY)-s.Content[p].ScaleY || s.Content[p].PositionY <= 0   {
    s.Content[p].SpeedY = -s.Content[p].SpeedY
  }
}


/*
Retourne vrai si la particle est sortie de l'écran de 5 fois sa taille
*/
func EstNonVisible(p Particle) bool {
  if p.PositionX+5*p.ScaleX >= float64(config.General.WindowSizeX) || p.PositionX <= -5*p.ScaleX || p.PositionY >= float64(config.General.WindowSizeY)+5*p.ScaleY || p.PositionY <= -5*p.ScaleY{
    return true
  }
  return false
}

















//
