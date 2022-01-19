package particles

// System définit un système de particules.
// Pour le moment il ne contient qu'un tableau de particules, mais cela peut
// évoluer durant votre projet.
//ajout de spawnrate permettant de respecter à la lettre le SpawnRate du json
type System struct {
	Content []Particle
	Spawnrate float64
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, et une
// opacité. Vous ajouterez certainement d'autres caractéristiques aux particules
// durant le projet.
//ajout de SpeedX et SpeedY qui définissent la vitesse de déplacement de la particules en X et en Y
//ajout de NonVisible qui permet de repérer les particules qui ne sont plus visible
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	SpeedX, SpeedY									float64
	NonVisible										  bool
	Vie 														int
}
