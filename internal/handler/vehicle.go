package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/rhinosc/code-review-1/internal"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
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
}

type BodyVehicleJSON struct {
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
}

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
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		var body BodyVehicleJSON
		err := request.JSON(r, &body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid body")
			return
		}

		//process

		// create vehicle
		vehicleAttributes := internal.VehicleAttributes{
			Brand:           body.Brand,
			Model:           body.Model,
			Registration:    body.Registration,
			Color:           body.Color,
			FabricationYear: body.FabricationYear,
			Capacity:        body.Capacity,
			MaxSpeed:        body.MaxSpeed,
			FuelType:        body.FuelType,
			Transmission:    body.Transmission,
			Weight:          body.Weight,
			Dimensions: internal.Dimensions{
				Height: body.Height,
				Length: body.Length,
				Width:  body.Width,
			},
		}
		vehicle := internal.Vehicle{
			VehicleAttributes: vehicleAttributes,
		}

		err = h.sv.Create(&vehicle)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "internal server error")
		}

		//response

		// serialize vehicle to JSON
		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles?color={color}&year={year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid year")
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.GetByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "vehicles not found")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByDimensions is a method that returns a handler for the route GET /vehicles/dimensions?length={min_length}-{max_length}&width={min_width}-{max_width}
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")

		lengthValues := strings.Split(length, "-")
		widthValues := strings.Split(width, "-")

		// fmt.Println(lengthValues, widthValues)
		minLength, err := strconv.ParseFloat(lengthValues[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid min length")
			return
		}
		maxLength, err := strconv.ParseFloat(lengthValues[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid max length")
			return
		}
		minWidth, err := strconv.ParseFloat(widthValues[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid min width")
			return
		}
		maxWidth, err := strconv.ParseFloat(widthValues[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid max width")
			return
		}
		// process
		// - get vehicles by dimensions
		v, err := h.sv.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			response.JSON(w, http.StatusNotFound, "vehicles not found")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeedByBrand is a method that returns a handler for the route GET /vehicles/average_speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get brand from URL
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.JSON(w, http.StatusBadRequest, "invalid brand")
			return
		}

		// process
		// - get average speed by brand
		averageSpeed, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.JSON(w, http.StatusNotFound, "vehicles not found")
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize average speed to JSON
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    averageSpeed,
		})
	}
}
