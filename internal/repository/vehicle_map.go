package repository

import (
	"fmt"

	"github.com/rhinosc/code-review-1/internal"
	"github.com/rhinosc/code-review-1/internal/loader"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(ld *loader.VehicleJSONFile, db map[int]internal.Vehicle, lastID int) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{ld: *ld, db: defaultDb, lastID: lastID}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// ld is the loader that loads the vehicles from a JSON file
	ld loader.VehicleJSONFile
	// db is a map of vehicles
	db     map[int]internal.Vehicle
	lastID int
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a vehicle
func (r *VehicleMap) Create(v *internal.Vehicle) (err error) {
	r.lastID++
	v.Id = r.lastID
	r.db[v.Id] = *v

	// save db to JSON file
	err = r.ld.Save(r.db)
	return
}

// GetByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter db
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = fmt.Errorf("no vehicles found with color %s and year %d", color, year)
		return
	}

	return
}

// GetByDimensions is a method that returns a map of vehicles by dimensions
func (r *VehicleMap) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// filter db
	for key, value := range r.db {
		if value.Length >= minLength && value.Length <= maxLength && value.Width >= minWidth && value.Width <= maxWidth {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = fmt.Errorf("no vehicles found with dimensions between %f and %f for length and between %f and %f for width", minLength, maxLength, minWidth, maxWidth)
		return
	}

	return
}
