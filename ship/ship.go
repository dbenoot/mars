package ship

type Spaceship struct {
	Name   string
	Water  int     // in liters
	Food   int     // in kilograms
	Fuel   int     // in liters
	Oxygen float64 // in %
	CO2    int     // in ppm
	Health int
}

type Location struct {
	Name        string
	Description string
	Transitions []string
	Subsystems  map[string]int
}

func New(name string, water int, food int, fuel int, oxygen float64, co2 int, health int) Spaceship {
	s := Spaceship{name, water, food, fuel, oxygen, co2, health}
	return s
}

func (s Spaceship) Process(lm map[string]Location, numastro int) Spaceship {

	// fuel expenditure based on engine efficiency
	s.Fuel = s.Fuel - ((100 / lm["Engineering"].Subsystems["engine"]) * 1)

	// oxygen expenditure based on #astronauts and oxygen reclamantion efficiency
	s.Oxygen = s.Oxygen - (float64(numastro) * 0.5) + (float64(lm["Engineering"].Subsystems["oxygen"]) / float64(100) * 1.5)
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
