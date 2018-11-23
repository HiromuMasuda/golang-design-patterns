package adapter

import (
	"fmt"
)

type (
	StudentA struct{}

	StudentB struct {
		studentA *StudentA
	}
)

func (studentA StudentA) enjoyWithClassmates() {
	fmt.Println("Lets Enjoy!")
}

func (studentB StudentB) organizeClass() {
	studentB.studentA.enjoyWithClassmates()
}

func RunDelegation() {
	studentA := &StudentA{}
	chairperson := &StudentB{studentA: studentA}
	chairperson.organizeClass()
}
