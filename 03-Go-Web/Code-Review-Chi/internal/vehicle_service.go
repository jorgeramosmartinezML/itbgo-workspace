package internal

import (
	"errors"
)

var (
	// ErrServiceVehicleAlreadyExists is an error that returns when a vehicle already exists
	ErrServiceVehicleAlreadyExists = errors.New("service: vehicle already exists")
	// ErrServiceVehicleInvalid is an error that returns when a vehicle is invalid
	ErrServiceVehicleInvalid = errors.New("service: invalid vehicle")
	// ErrServiceVehicleInternal is an error that returns when an internal error occurs
	ErrServiceVehicleInternal = errors.New("service: internal error")
	// ErrServiceRegistersNotFound is an error that returns when registers are not found
	ErrServiceRegistersNotFound = errors.New("service: registers not found")
	// ErrServicePeriodTimeInvalid is an error that returns when a period time is invalid
	ErrServicePeriodTimeInvalid = errors.New("service: period time invalid")
	// ErrServiceVehicleNotFound is an error that returns when a vehicle is not found
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")
	// ErrServiceUnknownVehicleProperty is an error that returns when a vehicle property is unknown
	ErrServiceUnknownVehicleProperty = errors.New("service: unknown vehicle property")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that creates a vehicle
	Create(p Vehicle) (err error)
	// FindAllByColorAndYear is a method that returns a map of vehicles by color and year
	FindAllByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// FindAllByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
	FindAllByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]Vehicle, err error)
	// GetAverageSpeedByBrand is a method that returns the average speed of a brand
	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
	// CreateAll is a method that creates a list of vehicles
	CreateAll(p []Vehicle) (err error)
	// Patch is a method that updates a vehicle
	Patch(id int, patch map[string]any) (v Vehicle, err error)
	// FindAllByFuelType is a method that returns a map of vehicles by fuel type
	FindAllByFuelType(fuelType string) (v map[int]Vehicle, err error)
	// DeleteById is a method that deletes a vehicle by id
	DeleteById(id int) (err error)
	// FindAllByTransmission is a method that returns a map of vehicles by transmission
	FindAllByTransmission(transmission string) (v map[int]Vehicle, err error)
	// GetAverageCapacityByBrand is a method that returns the average capacity of a brand
	GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error)
	// FindAllByDimensionsRange is a method that returns a map of vehicles by dimensions range
	FindAllByDimensionsRange(lengthFrom float64, lengthTo float64, widthFrom float64, widthTo float64) (v map[int]Vehicle, err error)
	// FindAllByWeightRange is a method that returns a map of vehicles by weight range
	FindAllByWeightRange(weightFrom float64, weightTo float64) (v map[int]Vehicle, err error)
}
