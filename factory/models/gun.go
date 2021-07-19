package models

type gun struct {
	Name  string
	Power int
}

func (g *gun) SetName(name string) {
	g.Name = name
}

func (g *gun) GetName() string {
	return g.Name
}

func (g *gun) SetPower(power int) {
	g.Power = power
}

func (g *gun) GetPower() int {
	return g.Power
}
