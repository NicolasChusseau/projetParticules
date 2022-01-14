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






















//
