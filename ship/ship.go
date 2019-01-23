package ship

type Spaceship struct {
	Name   string
	Water  int
	Food   int
	Fuel   int
	Oxygen int
	Health int
}

type Location struct {
	Name        string
	Description string
	Transitions []string
	Subsystems  []string
}

func New(name string, water int, food int, fuel int, oxygen int, health int) Spaceship {
	s := Spaceship{name, water, food, fuel, oxygen, health}
	return s
}

func (s Spaceship) Process() Spaceship {
	s.Fuel = s.Fuel - 1
	return s
}

func (s Spaceship) Eat(f int) (Spaceship, bool) {
	s.Food = s.Food - f
	if s.Food > 0 {
		return s, true
	} else {
		s.Food = 0
		return s, false
	}
}

func (s Spaceship) Drink(f int) (Spaceship, bool) {
	s.Water = s.Water - f
	if s.Water > 0 {
		return s, true
	} else {
		s.Water = 0
		return s, false
	}
}
