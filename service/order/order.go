package order

import (
	"firstAML/service/parcel"
	"fmt"
)

type Config struct {
	parcelConfig parcel.Config
}

func NewConfig(parcelConfig parcel.Config) *Config {
	config = &Config{parcelConfig}
	return config
}

var config *Config

type Order interface {
	TotalAmount() (float64, error)
	Description() string
}

func NewOrder() *baseOrder {
	return &baseOrder{}
}

func (o *baseOrder) AddParcel(parcelType, dimensions string, weight float64, shippedFrom, shippedTo string) error {
	if parcel := config.parcelConfig.Get(parcelType); parcel == nil {
		return fmt.Errorf("parcel type not recognized: %s", parcelType)
	}

	orderParcel := orderParcel{
		parcelType:  parcelType,
		dimensions:  dimensions,
		weight:      weight,
		shippedFrom: shippedFrom,
		shippedTo:   shippedTo,
	}
	o.parcels = append(o.parcels, orderParcel)
	return nil
}

type baseOrder struct {
	parcels []orderParcel
}

type orderParcel struct {
	parcelType  string
	dimensions  string
	weight      float64
	shippedFrom string
	shippedTo   string
	// TODO: add more logistics info
}

func (o baseOrder) TotalAmount() (float64, error) {
	var totalAmount float64
	for _, p := range o.parcels {
		parcel := config.parcelConfig.Get(p.parcelType)
		if parcel == nil {
			return 0, fmt.Errorf("parcel type not recognized: %s", p.parcelType)
		}
		totalAmount += parcel.BaseCost
	}

	return totalAmount, nil
}

func (o baseOrder) Description() string {
	var desc string
	for i, parcel := range o.parcels {
		if i == 0 {
			desc = fmt.Sprintf("there are %d parcels and these are details about them:", len(o.parcels))
		}
		desc = fmt.Sprintf("%s this is %d parcel, and its dimensions is %s, and its weight is %f, and its shipped from address is %s\n", desc, i, parcel.dimensions, parcel.weight, parcel.shippedFrom)
	}

	return desc
}

func ApplySpeedyShip(order Order) Order {
	return &speedyShip{order: order}
}

type speedyShip struct {
	cost  float64
	order Order
}

// speedy shipping doubles the cost of the entire order
func (s *speedyShip) TotalAmount() (float64, error) {
	cost, err := s.order.TotalAmount()
	if err != nil {
		return 0, fmt.Errorf("failed to get total amount %w", err)
	}
	s.cost = cost
	return cost + s.cost, nil
}

func (s *speedyShip) Description() string {
	return fmt.Sprintf("%s and it needs speedy ship service, and its cost is %f\n", s.order.Description(), s.cost)
}

func ApplyDiscount(order Order) Order {
	return &discount{order: order}
}

type discount struct {
	order Order
}

func (s *discount) TotalAmount() (float64, error) {
	return 0, nil
}

func (s *discount) Description() string {
	return ""
}
