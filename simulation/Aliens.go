package simulation

/*
	Alien store information and characterstics of alien. Going forward more infomation about the alien can be stored.
*/
type Alien struct {
	Name string
}

/*
	NewAlien simulates the creation/arrival of a new alien.
*/
func NewAlien(name string) *Alien {
	return &Alien{Name: name}
}
