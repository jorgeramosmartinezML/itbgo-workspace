package internal

import "errors"

var (
	// ErrRepositoryVehicleNotFound is an error that returns when a vehicle is not found
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")
	// ErrRepositoryRegistersNotFound is an error that returns when registers are not found
	ErrRepositoryRegistersNotFound = errors.New("repository: registers not found")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// CreateOrUpdate is a method that creates or updates a vehicle
	CreateOrUpdate(p Vehicle) (err error)
	// FindById is a method that returns a vehicle by id
	FindById(id int) (v Vehicle, err error)
	// FindAllByColorAndYear is a method that returns a map of vehicles by color and yearß
	FindAllByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// FindAllByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
	FindAllByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]Vehicle, err error)
	// GetAverageSpeedByBrand is a method that returns the average speed of a brand
	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
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
