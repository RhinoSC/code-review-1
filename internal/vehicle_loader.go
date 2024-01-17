package internal

// VehicleLoader is an interface that represents the loader for vehicles
type VehicleLoader interface {
	// Load is a method that loads the vehicles
	Load() (v map[int]Vehicle, err error)

	// Save is a method that saves the vehicles
	Save(v map[int]Vehicle) (err error)
}
