package service

import (
	"app/internal"
	"app/platform/patcher"
	"errors"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that creates a vehicle
func (s *VehicleDefault) Create(vehicle internal.Vehicle) error {
	_, err := s.rp.FindById(vehicle.Id)
	if err == nil {
		return internal.ErrServiceVehicleAlreadyExists
	}

	err = s.rp.CreateOrUpdate(vehicle)
	if err != nil {
		return internal.ErrServiceVehicleInternal
	}
	return nil
}

func (s *VehicleDefault) FindAllByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByColorAndYear(color, year)
	if len(v) == 0 {
		return nil, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) FindAllByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByBrandAndYearRange(brand, yearFrom, yearTo)
	if len(v) == 0 {
		return nil, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.GetAverageSpeedByBrand(brand)
	if errors.Is(err, internal.ErrRepositoryRegistersNotFound) {
		return 0, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) CreateAll(vehicles []internal.Vehicle) error {
	for _, vehicle := range vehicles {
		_, err := s.rp.FindById(vehicle.Id)
		if err == nil {
			return internal.ErrServiceVehicleAlreadyExists
		}
	}

	for _, vehicle := range vehicles {
		err := s.rp.CreateOrUpdate(vehicle)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *VehicleDefault) Patch(id int, patch map[string]any) (internal.Vehicle, error) {
	founded, err := s.rp.FindById(id)
	if err != nil {
		return founded, internal.ErrServiceVehicleNotFound
	}

	founded, err = patcher.Patch(founded, patch)
	if err != nil {
		fmt.Println(err.Error())
		return internal.Vehicle{}, internal.ErrServiceVehicleInvalid
	}

	err = s.rp.CreateOrUpdate(founded)
	if err != nil {
		return internal.Vehicle{}, internal.ErrServiceVehicleInternal
	}
	return founded, nil
}

func (s *VehicleDefault) FindAllByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByFuelType(fuelType)
	if len(v) == 0 {
		return v, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) DeleteById(id int) error {
	_, notExists := s.rp.FindById(id)
	if notExists != nil {
		return internal.ErrServiceVehicleNotFound
	}

	err := s.rp.DeleteById(id)
	return err
}

func (s *VehicleDefault) FindAllByTransmission(transmission string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByTransmission(transmission)
	if len(v) == 0 {
		return v, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	averageCapacity, err = s.rp.GetAverageCapacityByBrand(brand)
	if errors.Is(err, internal.ErrRepositoryRegistersNotFound) {
		return 0, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) FindAllByDimensionsRange(lengthFrom float64, lengthTo float64, widthFrom float64, widthTo float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByDimensionsRange(lengthFrom, lengthTo, widthFrom, widthTo)
	if len(v) == 0 {
		return v, internal.ErrServiceRegistersNotFound
	}
	return
}

func (s *VehicleDefault) FindAllByWeightRange(weightFrom float64, weightTo float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByWeightRange(weightFrom, weightTo)
	if len(v) == 0 {
		return v, internal.ErrServiceRegistersNotFound
	}
	return
}
