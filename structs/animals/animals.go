package animals

// Animal represents a multicellular, eukaryotic organism of the kingdom Animalia.
type Animal struct {
	Name string
	Hp   int
}

// Eater is something that eats other animal (aka a carnivore).
type Eater interface {
	Eat(Animal) bool
	Power() int
}

// Shark is a shark
type Shark struct {
	Animal
	BitingPower int
}

// Power determines the power of a shark's bite
func (s Shark) Power() int {
	return s.BitingPower
}

// Eat determines if a shark can eat another animal.
func (s Shark) Eat(a Animal) bool {
	return s.Power() > a.Hp
}
