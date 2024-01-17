package internal

import "errors"

var (
	// ErrVehicleNotFound is an error that represents a vehicle not found
	ErrVehicleNotFound = errors.New("vehicle not found")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// Create is a method that creates a vehicle
	Create(v *Vehicle) (err error)

	// GetByColorAndYear is a method that returns a map of vehicles by color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	// GetByDimensions is a method that returns a map of vehicles by dimensions
	GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]Vehicle, err error)

	// GetAverageSpeedByBrand is a method that returns the average speed of a vehicle
	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
}
