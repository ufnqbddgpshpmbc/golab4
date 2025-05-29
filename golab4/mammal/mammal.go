package mammal

type Mammal struct {
	Name    string
	Age     int
	Species string
}

func NewMammal(name string, age int, species string) Mammal {
	return Mammal{
		Name:    name,
		Age:     age,
		Species: species,
	}
}

func SetMammalName(mammal *Mammal, name string) {
	mammal.Name = name
}

func SetMammalAge(mammal *Mammal, age int) {
	mammal.Age = age
}

func SetMammalSpecies(mammal *Mammal, species string) {
	mammal.Species = species
}

func PrintMammal(mammal Mammal) {
	println("Name:", mammal.Name)
	println("Age:", mammal.Age)
	println("Species:", mammal.Species)
}

func WipeMammal(mammal *Mammal) {
	mammal.Name = ""
	mammal.Age = 0
	mammal.Species = ""
}
