package model

import "app/internal"

func ToVehicleJSON(v internal.Vehicle) (vehicleJSON VehicleJSON) {
	vehicleJSON = VehicleJSON{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Height,
		Length:          v.Length,
		Width:           v.Width,
	}
	return vehicleJSON
}

func ToVehicle(v VehicleJSON) (vehicle internal.Vehicle) {
	vehicle = internal.Vehicle{
		Id: v.ID,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Dimensions: internal.Dimensions{
				Height: v.Height,
				Length: v.Length,
				Width:  v.Width,
			},
		},
	}
	return vehicle
}
