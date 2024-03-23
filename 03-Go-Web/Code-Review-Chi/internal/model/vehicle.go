package model

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id" validate:"required,gt=0"`
	Brand           string  `json:"brand" validate:"required,gte=1"`
	Model           string  `json:"model" validate:"required,gte=1"`
	Registration    string  `json:"registration" validate:"required,gte=1"`
	Color           string  `json:"color" validate:"required,gte=1"`
	FabricationYear int     `json:"year" validate:"required,gte=1560"`
	Capacity        int     `json:"passengers" validate:"required,gte=1"`
	MaxSpeed        float64 `json:"max_speed" validate:"required,gte=1"`
	FuelType        string  `json:"fuel_type" validate:"required,gte=1"`
	Transmission    string  `json:"transmission" validate:"required,gte=1"`
	Weight          float64 `json:"weight" validate:"required,gte=0"`
	Height          float64 `json:"height" validate:"required,gte=0"`
	Length          float64 `json:"length" validate:"gte=0"`
	Width           float64 `json:"width" validate:"required,gte=0"`
}
