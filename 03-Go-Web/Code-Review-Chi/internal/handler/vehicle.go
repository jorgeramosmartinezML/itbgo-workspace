package handler

import (
	"app/internal"
	"app/internal/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
/*type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}*/

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vehicle model.VehicleJSON
		err := request.JSON(r, &vehicle)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		validate := validator.New()
		err = validate.Struct(vehicle)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body by wrong or missing data")
			return
		}

		err = h.sv.Create(model.ToVehicle(vehicle))
		if err != nil {
			h.handleError(err, w)
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{"message": "product created"})
	}
}

func (h *VehicleDefault) GetAllByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid year")
			return
		}

		// process
		// - get all vehicles by color and year
		v, err := h.sv.FindAllByColorAndYear(color, year)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetAllByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		brand := chi.URLParam(r, "brand")
		yearFrom, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid start year")
			return
		}
		yearTo, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid end year")
			return
		}

		// process
		// - get all vehicles by brand and year range
		v, err := h.sv.FindAllByBrandAndYearRange(brand, yearFrom, yearTo)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		brand := chi.URLParam(r, "brand")

		// process
		// - get average speed by brand
		v, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

func (h *VehicleDefault) CreateAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		var vehiclesJson []model.VehicleJSON
		if err := json.NewDecoder(r.Body).Decode(&vehiclesJson); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		validate := validator.New()
		for _, vehicle := range vehiclesJson {
			err := validate.Struct(vehicle)
			if err != nil {
				response.Error(w, http.StatusBadRequest, "Invalid request body by wrong or missing data")
				return
			}
		}

		var vehicles []internal.Vehicle
		for _, vehicle := range vehiclesJson {
			vehicles = append(vehicles, model.ToVehicle(vehicle))
		}

		// process
		// - create all vehicles
		err := h.sv.CreateAll(vehicles)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{"message": "products created"})
	}
}

func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
		}

		type PatchBody struct {
			MaxSpeed float64 `json:"max_speed"`
		}

		var body PatchBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body by wrong or missing data")
			return
		}
		bodyMap := map[string]any{"max_speed": body.MaxSpeed}

		maxSpeedFloat := bodyMap["max_speed"].(float64)
		if maxSpeedFloat <= 0 {
			response.Error(w, http.StatusBadRequest, "Invalid request body by wrong max_speed")
			return
		}

		// process
		updated, err := h.sv.Patch(id, bodyMap)
		if err != nil {
			h.handleError(err, w)
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    model.ToVehicleJSON(updated),
		})

	}
}

func (h *VehicleDefault) FindAllByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		fuelType := chi.URLParam(r, "type")

		// process
		// - get all vehicles by fuel type
		v, err := h.sv.FindAllByFuelType(fuelType)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		// process
		// - delete vehicle by id
		err = h.sv.DeleteById(id)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, map[string]any{"message": "success"})
	}
}

func (h *VehicleDefault) FindAllByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		transmission := chi.URLParam(r, "type")

		// process
		// - get all vehicles by transmission
		v, err := h.sv.FindAllByTransmission(transmission)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		type PatchBody struct {
			FuelType string `json:"fuel_type"`
		}

		var body PatchBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body by wrong or missing data")
			return
		}
		bodyMap := map[string]any{"fuel_type": body.FuelType}

		// process
		updated, err := h.sv.Patch(id, bodyMap)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    model.ToVehicleJSON(updated),
		})
	}
}

func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		brand := chi.URLParam(r, "brand")

		// process
		// - get average capacity by brand
		average, err := h.sv.GetAverageCapacityByBrand(brand)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    average,
		})
	}
}

func (h *VehicleDefault) FindAllByDimensionsRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")

		partsLength := strings.Split(length, "-")
		if len(partsLength) != 2 {
			response.Error(w, http.StatusBadRequest, "Invalid length")
			return
		}
		partsWidth := strings.Split(width, "-")
		if len(partsWidth) != 2 {
			response.Error(w, http.StatusBadRequest, "Invalid width")
			return
		}
		minLength, err := strconv.ParseFloat(partsLength[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid length")
			return
		}
		maxLength, err := strconv.ParseFloat(partsLength[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid length")
			return
		}
		minWidth, err := strconv.ParseFloat(partsWidth[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid width")
			return
		}
		maxWidth, err := strconv.ParseFloat(partsWidth[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid width")
			return
		}

		// process
		// - get all vehicles by dimensions range
		v, err := h.sv.FindAllByDimensionsRange(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) FindAllByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		weightMinQuery := r.URL.Query().Get("min")
		weightMaxQuery := r.URL.Query().Get("max")

		minWeight, err := strconv.ParseFloat(weightMinQuery, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid min weight")
			return
		}
		maxWeight, err := strconv.ParseFloat(weightMaxQuery, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid max weight")
			return
		}

		// process
		// - get all vehicles by weight range
		v, err := h.sv.FindAllByWeightRange(minWeight, maxWeight)
		if err != nil {
			h.handleError(err, w)
			return
		}

		// response
		data := make(map[int]model.VehicleJSON)
		for key, value := range v {
			data[key] = model.ToVehicleJSON(value)
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) handleError(err error, w http.ResponseWriter) {
	switch err {
	case internal.ErrServiceVehicleAlreadyExists:
		response.Error(w, http.StatusConflict, "Vehicle already exists")
	case internal.ErrServiceVehicleInvalid:
		response.Error(w, http.StatusBadRequest, "Invalid vehicle by wrong data or missing data")
	case internal.ErrServicePeriodTimeInvalid:
		response.Error(w, http.StatusBadRequest, "Invalid period time")
	case internal.ErrServiceRegistersNotFound:
		response.Error(w, http.StatusNotFound, "Registers not found")
	case internal.ErrServiceVehicleNotFound:
		response.Error(w, http.StatusNotFound, "Vehicle not found")
	default:
		response.Error(w, http.StatusInternalServerError, "Internal server error")
	}
}
