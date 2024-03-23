package patcher

import (
	"app/internal"
	"errors"
	"strings"
)

var (
	// ErrPatcherInvalidType is an error that returns when the patcher receives an invalid type
	ErrPatcherInvalidType = errors.New("patcher: invalid type")
	// ErrPatcherUnknownVehicleProperty is an error that returns when a vehicle property is unknown
	ErrPatcherUnknownVehicleProperty = errors.New("service: unknown vehicle property")
)

func Patch(vehicle internal.Vehicle, patchBody map[string]interface{}) (internal.Vehicle, error) {
	for key, value := range patchBody {
		switch strings.ToLower(key) {
		case "brand":
			vehicle.Brand = value.(string)
		case "model":
			vehicle.Model = value.(string)
		case "registration":
			vehicle.Registration = value.(string)
		case "color":
			vehicle.Color = value.(string)
		case "year":
			vehicle.FabricationYear = value.(int)
		case "passengers":
			vehicle.Capacity = value.(int)
		case "max_speed":
			vehicle.MaxSpeed = value.(float64)
		case "fuel_type":
			vehicle.FuelType = value.(string)
		case "transmission":
			vehicle.Transmission = value.(string)
		case "weight":
			vehicle.Weight = value.(float64)
		case "height":
			vehicle.Height = value.(float64)
		case "length":
			vehicle.Length = value.(float64)
		case "width":
			vehicle.Width = value.(float64)
		default:
			return internal.Vehicle{}, internal.ErrServiceUnknownVehicleProperty
		}
	}
	return vehicle, nil
}
