package world

type Alien struct {
	Name        string
	CurrentCity *City
}

func NewAlien(name string) *Alien {
	return &Alien{Name: name}
}
