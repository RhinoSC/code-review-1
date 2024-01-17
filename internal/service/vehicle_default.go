package service

import (
	"fmt"

	"github.com/rhinosc/code-review-1/internal"
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
func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

// GetByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	if err != nil {
		err = fmt.Errorf("error getting vehicles by color and year: %w", err)
	}
	return
}

// GetByDimensions is a method that returns a map of vehicles by dimensions
func (s *VehicleDefault) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
	if err != nil {
		err = fmt.Errorf("error getting vehicles by dimensions: %w", err)
	}
	return
}

// GetAverageSpeedByBrand is a method that returns the average speed of a vehicle
func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.GetAverageSpeedByBrand(brand)
	if err != nil {
		err = fmt.Errorf("error getting average speed by brand: %w", internal.ErrVehicleNotFound)
	}
	return
}
