package particles
import (
  "testing"
  "project-particles/config"
  "log"
)


func TestNewSystem1(t *testing.T)  {
  log.Println("a")
  config.General.InitNumParticles = 5
  s := NewSystem()
  if len(s.Content) != config.General.InitNumParticles{
    t.Fail()
  }
}

func TestNewSystem2(t *testing.T)  {
  config.General.InitNumParticles = 5
  s := NewSystem()
  if !config.General.RandomSpawn{
    for i := 1; i < len(s.Content); i++ {
      if s.Content[i-1].PositionX != s.Content[i].PositionX || s.Content[i-1].PositionY != s.Content[i].PositionY{
        t.Fail()
      }
    }
  }
}

func TestUpdate1(t *testing.T)  {
  config.General.InitNumParticles = 5
  config.General.WindowSizeX = 100
  config.General.WindowSizeY = 100
  config.General.SpawnX = 50
  config.General.SpawnY = 50
  config.General.MaxParticles=6513204651320
  s := NewSystem()
  config.General.SpawnRate = 1
  s.Update()
  if float64(len(s.Content)) != 6{
    t.Fail()
  }
  s.Update()
  s.Update()
  s.Update()
  s.Update()
  if float64(len(s.Content)) != 10{
    t.Fail()
  }
}


func TestUpdate2(t *testing.T)  {
  s := NewSystem()
  config.General.SpawnRate = 2.5
  config.General.MaxParticles=6513204651320
  s.Update()
  if len(s.Content) != 7{
    t.Fail()
  }
  s.Update()
  s.Update()
  s.Update()
  s.Update()
  if float64(len(s.Content)) != 17{
    t.Fail()
  }
}


func TestUpdate3(t* testing.T)  {
  s := NewSystem()
  config.General.MaxParticles=6513204651320
  posX := s.Content[0].PositionX
  posY := s.Content[0].PositionY
  s.Update()
  if posX+s.Content[0].SpeedX != s.Content[0].PositionX || posY+s.Content[0].SpeedY != s.Content[0].PositionY{
    t.Fail()
  }
}

func TestMaxParticles(t *testing.T)  {
  config.General.InitNumParticles=50
  config.General.WindowSizeX=1000
  config.General.WindowSizeY=1000
  config.General.SpawnRate=10
  config.General.MaxParticles=60
  s := NewSystem()
  s.Update()
  if len(s.Content) != 60{
    t.Fail()
  }
  s.Update()
  if len(s.Content) != 60{
    t.Fail()
  }
  s.Update()
  if len(s.Content) != 60{
    t.Fail()
  }
  s.Update()
  if len(s.Content) != 60{
    t.Fail()
  }
}

func TestEstNonVisible(t *testing.T)  {
  config.General.WindowSizeX=5
  config.General.WindowSizeY=5
  p := Particle{
    PositionX: 2.5,
    PositionY: 2.5,
    NonVisible: false,
  }
  if EstNonVisible(p){
    t.Fail()
  }
  p.PositionX = 10
  if !EstNonVisible(p){
    t.Fail()
  }
}

func TestNewSystemplouf(t *testing.T)  {
  config.General.Configuration="plouf"
  config.General.InitNumParticles = 5
  config.General.WindowSizeX = 800
  config.General.WindowSizeY = 600
  particle_eau := 1800
  s := NewSystem()
  if len(s.Content) != config.General.InitNumParticles+particle_eau{
    log.Println("a",len(s.Content))
    log.Println("b", particle_eau)
    t.Fail()
  }
}

func TestUpdateplouf1(t *testing.T)  {
  config.General.Configuration="plouf"
  config.General.InitNumParticles = 5
  config.General.WindowSizeX = 100
  config.General.WindowSizeY = 100
  config.General.SpawnX = 50
  config.General.SpawnY = 50
  config.General.MaxParticles=6513204651320
  s := NewSystem()
  var particule_bleu int = 2400
  config.General.SpawnRate = 1
  s.Update()
  if len(s.Content) != 6+particule_bleu{
    t.Fail()
  }
  s.Update()
  s.Update()
  s.Update()
  s.Update()
  if len(s.Content) != 10+particule_bleu{
    t.Fail()
  }
}

func TestUpdateplouf2(t *testing.T)  {
  config.General.Configuration="plouf"
  s := NewSystem()
  var particule_bleu int = 2400
  config.General.SpawnRate = 2.5
  config.General.MaxParticles=6513204651320
  s.Update()
  if len(s.Content) != 7+particule_bleu{
    t.Fail()
  }
  s.Update()
  s.Update()
  s.Update()
  s.Update()
  if len(s.Content) != 17+particule_bleu{
    t.Fail()
  }
}

func TestUpdateplouf3(t* testing.T)  {
  config.General.Configuration="plouf"
  s := NewSystem()
  config.General.MaxParticles=6513204651320
  posX := s.Content[0].PositionX
  posY := s.Content[0].PositionY
  s.Update()
  if posX+s.Content[0].SpeedX != s.Content[0].PositionX || posY+s.Content[0].SpeedY != s.Content[0].PositionY{
    t.Fail()
  }
}

func TestMaxParticlesplouf(t *testing.T)  {
  config.General.Configuration="plouf"
  config.General.InitNumParticles=50
  config.General.WindowSizeX=100
  config.General.WindowSizeY=100
  config.General.SpawnRate=10
  var particule_bleu int = 2400
  config.General.MaxParticles=particule_bleu+80
  s := NewSystem()
  s.Update()
  if len(s.Content) != particule_bleu+60{
    t.Fail()
  }
  s.Update()
  if len(s.Content) != particule_bleu+70{
    t.Fail()
  }
  s.Update()
  if len(s.Content) != particule_bleu+80{
    t.Fail()
  }
  s.Update()
  if len(s.Content) != particule_bleu+80{
    t.Fail()
  }
}

func TestVagueplouf1(t *testing.T)  {
  config.General.Configuration="plouf"
  config.General.WindowSizeX=800
  config.General.WindowSizeY=600
  config.General.MaxParticles=9999
  config.General.InitNumParticles=1
  s := NewSystem()
  s.Content[18].Vague=10
  s.Update()
  if s.Content[18].Vague != 1 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 2 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 3 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 4 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 0 {
    t.Fail()
  }
}

func TestVagueplouf2(t *testing.T)  {
  config.General.Configuration="plouf"
  config.General.WindowSizeX=800
  config.General.WindowSizeY=600
  config.General.MaxParticles=9999
  config.General.InitNumParticles=1
  s := NewSystem()
  s.Content[18].Vague=9
  s.Update()
  if s.Content[18].Vague != 5 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 6 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 7 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 8 {
    t.Fail()
  }
  s.Update()
  if s.Content[18].Vague != 0 {
    t.Fail()
  }
}

func TestEstDansLo(t *testing.T)  {
  config.General.Configuration="plouf"
  config.General.WindowSizeX = 800
  config.General.WindowSizeY = 600
  config.General.MaxParticles = 666666
  s := NewSystem()
  p := Particle{
    ScaleX:    3, ScaleY: 1,
    PositionX: 400,
    PositionY: 500,
    ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
    Opacity: 1,
    SpeedX: 0, SpeedY: 0,
  }
  s.Content = append(s.Content,p)
  if !EstDansLo(s.Content[len(s.Content)-1]) {
    t.Fail()
  }
  p = Particle{
    ScaleX:    3, ScaleY: 1,
    PositionX: 400,
    PositionY: 100,
    ColorRed: 0.5, ColorGreen: 0.5, ColorBlue: 0.5,
    Opacity: 1,
    SpeedX: 0, SpeedY: 0,
  }
  s.Content = append(s.Content,p)
  if EstDansLo(s.Content[len(s.Content)-1]) {
    t.Fail()
  }
}


func TestNewSystemneige1(t *testing.T)  {
  config.General.Configuration="neige"
  config.General.InitNumParticles = 5
  s := NewSystem()
  if len(s.Content) != config.General.InitNumParticles{
    t.Fail()
  }
}

func TestNewSystemneige2(t *testing.T)  {
  config.General.Configuration="neige"
  config.General.InitNumParticles = 5
  s := NewSystem()
  if !config.General.RandomSpawn{
    for i := 1; i < len(s.Content); i++ {
      if s.Content[i-1].PositionX != s.Content[i].PositionX || s.Content[i-1].PositionY != s.Content[i].PositionY{
        t.Fail()
      }
    }
  }
}


















//
