package human

import (
	"example.com/mammal"
)

type Human struct {
	mammal.Mammal
	Language string
	Weight   int
	Height   int
}

func NewHuman(mammal mammal.Mammal, language string, weight int, height int) Human {
	return Human{
		Mammal:   mammal,
		Language: language,
		Weight:   weight,
		Height:   height,
	}
}

func SetHumanLanguage(human *Human, language string) {
	human.Language = language
}

func SetHumanWeight(human *Human, weight int) {
	human.Weight = weight
}

func SetHumanHeight(human *Human, height int) {
	human.Height = height
}

func PrintHuman(human Human) {
	mammal.PrintMammal(human.Mammal)
	println("Language:", human.Language)
	println("Weight:", human.Weight)
	println("Height:", human.Height)
}

func WipeHuman(human *Human) {
	mammal.WipeMammal(&human.Mammal)
	human.Language = ""
	human.Weight = 0
	human.Height = 0
}
