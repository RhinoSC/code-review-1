package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rhinosc/code-review-1/internal"
)

// NewVehicleJSONFile is a function that returns a new instance of VehicleJSONFile
func NewVehicleJSONFile(path string) *VehicleJSONFile {
	return &VehicleJSONFile{
		path: path,
	}
}

// VehicleJSONFile is a struct that implements the LoaderVehicle interface
type VehicleJSONFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	Id              int     `json:"id"`
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

// Load is a method that loads the vehicles
func (l *VehicleJSONFile) Load() (v map[int]internal.Vehicle, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var vehiclesJSON []VehicleJSON
	err = json.NewDecoder(file).Decode(&vehiclesJSON)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]internal.Vehicle)
	for _, vh := range vehiclesJSON {
		v[vh.Id] = internal.Vehicle{
			Id: vh.Id,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           vh.Brand,
				Model:           vh.Model,
				Registration:    vh.Registration,
				Color:           vh.Color,
				FabricationYear: vh.FabricationYear,
				Capacity:        vh.Capacity,
				MaxSpeed:        vh.MaxSpeed,
				FuelType:        vh.FuelType,
				Transmission:    vh.Transmission,
				Weight:          vh.Weight,
				Dimensions: internal.Dimensions{
					Height: vh.Height,
					Length: vh.Length,
					Width:  vh.Width,
				},
			},
		}
	}

	return
}

// Save is a method that saves the vehicles
func (l *VehicleJSONFile) Save(v map[int]internal.Vehicle) (err error) {
	f, err := os.OpenFile(l.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var vehiclesJSON []VehicleJSON
	for _, vh := range v {
		vehiclesJSON = append(vehiclesJSON, VehicleJSON{
			Id:              vh.Id,
			Brand:           vh.Brand,
			Model:           vh.Model,
			Registration:    vh.Registration,
			Color:           vh.Color,
			FabricationYear: vh.FabricationYear,
			Capacity:        vh.Capacity,
			MaxSpeed:        vh.MaxSpeed,
			FuelType:        vh.FuelType,
			Transmission:    vh.Transmission,
			Weight:          vh.Weight,
			Height:          vh.Height,
			Length:          vh.Length,
			Width:           vh.Width,
		})
	}

	err = json.NewEncoder(f).Encode(vehiclesJSON)
	if err != nil {
		fmt.Println("error encoding file: ", l.path)
		return
	}
	return
}
