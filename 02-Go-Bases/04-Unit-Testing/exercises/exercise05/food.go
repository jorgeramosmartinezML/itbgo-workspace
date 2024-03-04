package exercise05

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func DogFood(amount int) float64 {
	return float64(amount) * 10
}

func CatFood(amount int) float64 {
	return float64(amount) * 5
}

func HamsterFood(amount int) float64 {
	return float64(amount) * 0.25
}

func TarantulaFood(amount int) float64 {
	return float64(amount) * 0.15
}

func animal(animal string) (func(int) float64, string) {
	switch animal {
	case Dog:
		return DogFood, ""
	case Cat:
		return CatFood, ""
	case Hamster:
		return HamsterFood, ""
	case Tarantula:
		return TarantulaFood, ""
	default:
		return nil, "Animal not found"
	}
}
