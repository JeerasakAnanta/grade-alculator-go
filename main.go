package main

import "fmt"

// ----- Interface Definitions -----
type GradeCalculator interface {
	Calculate(score float64) string
}

type Printer interface {
	Print(name string, score float64, grade string)
}

// ----- Concrete Implementations -----
type StandardGradeCalculator struct{}

func (s StandardGradeCalculator) Calculate(score float64) string {
	switch {
	case score >= 80:
		return "A"
	case score >= 70:
		return "B"
	case score >= 60:
		return "C"
	case score >= 50:
		return "D"
	default:
		return "F"
	}
}

type ConsolePrinter struct{}

func (c ConsolePrinter) Print(name string, score float64, grade string) {
	fmt.Printf("Student: %s | Score: %.2f | Grade: %s\n", name, score, grade)
}

// ----- High-Level Module -----
type GradeService struct {
	calculator GradeCalculator
	printer    Printer
}

func NewGradeService(c GradeCalculator, p Printer) *GradeService {
	return &GradeService{calculator: c, printer: p}
}

func (g *GradeService) Process(name string, score float64) {
	grade := g.calculator.Calculate(score)
	g.printer.Print(name, score, grade)
}

// ----- Main Function -----
func main() {

	calculator := StandardGradeCalculator{}
	printer := ConsolePrinter{}
	service := NewGradeService(calculator, printer)

	//
	fmt.Println("===============================================")
	fmt.Println("|            Grade Calculator                 |")
	fmt.Println("===============================================")

	service.Process("Jeerasak", 85)
	service.Process("Ananta", 75)
	service.Process("Game", 65)
	service.Process("Jo", 65)
	service.Process("Pong", 45)

	fmt.Println("===============================================")

}
