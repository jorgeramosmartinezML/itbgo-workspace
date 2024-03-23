package repository

import (
	"app/internal"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (repo *VehicleMap) FindAll() (vehicle map[int]internal.Vehicle, err error) {
	vehicle = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range repo.db {
		vehicle[key] = value
	}

	return
}

// Create is a method that creates a vehicle
func (repo *VehicleMap) CreateOrUpdate(vehicle internal.Vehicle) (err error) {
	// add vehicle to db
	repo.db[vehicle.Id] = vehicle
	return
}

// FindById is a method that returns a vehicle by id
func (r *VehicleMap) FindById(id int) (v internal.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}
	return
}

func (r *VehicleMap) FindAllByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) FindAllByBrandAndYearRange(brand string, yearFrom int, yearTo int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= yearFrom && value.FabricationYear <= yearTo {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed = 0.0
	vehicles := 0

	for _, value := range r.db {
		if value.Brand == brand {
			averageSpeed += value.MaxSpeed
			vehicles++
		}
	}

	if vehicles == 0 {
		err = internal.ErrRepositoryRegistersNotFound
		return
	}

	averageSpeed /= float64(vehicles)
	return
}

func (r *VehicleMap) FindAllByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.FuelType == fuelType {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) DeleteById(id int) (err error) {
	delete(r.db, id)
	return
}

func (r *VehicleMap) FindAllByTransmission(transmission string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Transmission == transmission {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	totalCapacity := 0
	vehicles := 0

	for _, value := range r.db {
		if value.Brand == brand {
			totalCapacity += value.Capacity
			vehicles++
		}
	}

	if vehicles == 0 {
		err = internal.ErrRepositoryRegistersNotFound
		return
	}

	averageCapacity = float64(totalCapacity) / float64(vehicles)
	return
}

func (r *VehicleMap) FindAllByDimensionsRange(lengthFrom float64, lengthTo float64, widthFrom float64, widthTo float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Length >= lengthFrom && value.Length <= lengthTo && value.Width >= widthFrom && value.Width <= widthTo {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindAllByWeightRange(weightFrom float64, weightTo float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Weight >= weightFrom && value.Weight <= weightTo {
			v[key] = value
		}
	}
	return
}
