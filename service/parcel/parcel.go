package parcel

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	parcels []parcel `yaml:"parcels"`
}

func Init(filePath string) (*Config, error) {
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %w", err)
	}

	fmt.Println("config info from file:", string(configFile))

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml file %w", err)
	}
	fmt.Printf("parcel config is created: %+v \n", config)
	return &config, nil
}

func (c Config) Get(parcelType string) *parcel {
	for _, p := range c.parcels {
		if p.Type == parcelType {
			return &p
		}
	}

	return nil
}

// TODO: needs to think about required/optional fields
func AddNewParcelType(parcelType, dimensions, weightLimit *string, baseCost, costPerKiloOverLimit *float64) (string, error) {
	var newParcel parcel
	if parcelType != nil {
		newParcel.Type = *parcelType
	}
	if dimensions != nil {
		newParcel.Dimensions = *dimensions
	}
	if weightLimit != nil {
		newParcel.WeightLimit = *weightLimit
	}
	if baseCost != nil {
		newParcel.BaseCost = *baseCost
	}
	if costPerKiloOverLimit != nil {
		newParcel.CostPerKiloOverLimit = *costPerKiloOverLimit
	}
	return newParcel.Description(), nil
}

type parcel struct {
	Type                 string  `yaml:"type"`
	Dimensions           string  `yaml:"dimensions"`
	BaseCost             float64 `yaml:"baseCost"`
	WeightLimit          string  `yaml:"weightLimit"`
	CostPerKiloOverLimit float64 `yaml:"costPerKiloOverLimit"`
}

// TODO: how to remove empty fields in description?
func (p parcel) Description() string {
	return fmt.Sprintf("this is %s parcel, and its dimension is %s, and its cost is %f, and its weight limit is %s, and its cost per kilo over weight is %f", p.Type, p.Dimensions, p.BaseCost, p.WeightLimit, p.CostPerKiloOverLimit)
}
