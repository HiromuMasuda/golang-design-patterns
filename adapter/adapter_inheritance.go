package adapter

import (
	"fmt"
)

type StudentC struct{}

func (studentC StudentC) enjoyWithClassmates() {
	fmt.Println("Lets Enjoy!")
}

func (studentC StudentC) organizeClass() {
	studentC.enjoyWithClassmates()
}

func RunInheritance() {
	chairperson := &StudentC{}
	chairperson.organizeClass()
}
