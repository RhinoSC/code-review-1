package repository

import (
	"github.com/rhinosc/code-review-1/internal"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle, lastID int) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb, lastID: lastID}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
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
	return
}
