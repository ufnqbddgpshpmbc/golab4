package docent

import (
	"github.com/ufnqbddgpshpmbc/golab4/human"
)

type Docent struct {
	human.Human
	Department string
	Experience int
	Salary     float64
}

func NewDocent(human human.Human, department string, experience int, salary float64) Docent {
	return Docent{
		Human:      human,
		Department: department,
		Experience: experience,
		Salary:     salary,
	}
}

func SetDocentDepartment(docent *Docent, department string) {
	docent.Department = department
}

func SetDocentExperience(docent *Docent, experience int) {
	docent.Experience = experience
}

func SetDocentSalary(docent *Docent, salary float64) {
	docent.Salary = salary
}

func PrintDocent(docent Docent) {
	human.PrintHuman(docent.Human)
	println("Department:", docent.Department)
	println("Experience:", docent.Experience)
	println("Salary:", docent.Salary)
}

func WipeDocent(docent *Docent) {
	human.WipeHuman(&docent.Human)
	docent.Department = ""
	docent.Experience = 0
	docent.Salary = 0.0
}
