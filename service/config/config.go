package config

import (
	"firstAML/service/order"
	"firstAML/service/parcel"
	"fmt"
)

func NewConfig(parcelConfigFilePath string) error {
	parcelConfig, err := parcel.Init(parcelConfigFilePath)
	if err != nil {
		return fmt.Errorf("failed to init parcel config: %w", err)
	}

	order.NewConfig(*parcelConfig)
	return nil
}
