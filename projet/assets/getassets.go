package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"project-particles/config"
)

// ParticleImage est une variable globale pour stocker l'image d'une particule
var ParticleImage *ebiten.Image

// Get charge en mémoire l'image de la particule. (particle.png)
// Vous pouvez changer cette image si vous le souhaitez, et même en proposer
// plusieurs, qu'on peut choisir via le fichier de configuration. Cependant
// ceci n'est pas du tout central dans le projet et ne devrait être fait que
// si vous avez déjà bien avancé sur tout le reste.
func Get() {
	var err error
	ParticleImage, _, err = ebitenutil.NewImageFromFile(config.General.ParticleImage)
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
}
