package template_method

import "fmt"

type (
	Wood struct {
		Name string
	}

	WoodCutPrint interface {
		Drow(w *Wood)
		Cut(w *Wood)
		Print(w *Wood)
	}

	WoodCutPrintA struct {
		Wood *Wood
	}
	WoodCutPrintB struct {
		Wood *Wood
	}
)

func (w WoodCutPrintA) Drow() {
	fmt.Printf("%s drow in A style\n", w.Wood.Name)
}

func (w WoodCutPrintA) Cut() {
	fmt.Printf("%s cut in A style\n", w.Wood.Name)
}

func (w WoodCutPrintA) Print() {
	fmt.Printf("%s print in A style\n", w.Wood.Name)
}

func (w WoodCutPrintB) Drow() {
	fmt.Printf("%s drow in B style\n", w.Wood.Name)
}

func (w WoodCutPrintB) Cut() {
	fmt.Printf("%s cut in B style\n", w.Wood.Name)
}

func (w WoodCutPrintB) Print() {
	fmt.Printf("%s print in B style\n", w.Wood.Name)
}

func Run() {
	woodA := &Wood{Name: "woodA"}
	woodB := &Wood{Name: "woodB"}
	woodCutPrintA := &WoodCutPrintA{Wood: woodA}
	woodCutPrintB := &WoodCutPrintB{Wood: woodB}

	woodCutPrintA.Drow()
	woodCutPrintA.Cut()
	woodCutPrintA.Print()

	woodCutPrintB.Drow()
	woodCutPrintB.Cut()
	woodCutPrintB.Print()
}
