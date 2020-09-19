package patterns

import "fmt"

type Progress interface {
	Progress()
}

type Progression struct {
	first    int
	last     int
	H        int
	progList []int
}

func (p *Progression) TemplateMethod() {
	p.InitializeProgression(p.first, p.last, p.H)
}

func (p *Progression) InitializeProgression(f, l, h int) {
	p.first = f
	p.last = l
	p.H = h
}

func (p Progression) Print() {
	fmt.Println("Последовательность:")
	for _, item := range p.progList {
		fmt.Printf(" %v", item)
	}
	fmt.Println()
}

type ArithmeticProgression struct {
	Progression
}

func NewArithmeticProgression(f, l, h int) ArithmeticProgression {
	return ArithmeticProgression{Progression{first: f, last: l, H: h}}
}

func (p *ArithmeticProgression) Progress() {
	fF := p.first

	for ok := true; ok; ok = fF < p.last {
		p.progList = append(p.progList, fF)
		fF += p.H
	}
}

type GeometryProgression struct {
	Progression
}

func NewGeometryProgression(f, l, h int) GeometryProgression {
	return GeometryProgression{Progression{first: f, last: l, H: h}}
}

func (p *GeometryProgression) Progress() {
	fF := p.first

	for ok := true; ok; ok = fF < p.last {
		p.progList = append(p.progList, fF)
		fF *= p.H
	}
}
