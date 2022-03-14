package main

import (
	"fmt"
	"log"

	"firstAML/service/config"
	"firstAML/service/order"
)

func main() {
	config := config.NewConfig("./service/parcel/config.yaml")
	fmt.Printf("starting: %+v", config)
	o := order.NewOrder()
	if err := o.AddParcel("small", "2cm", 2, "nelson", "auckland"); err != nil {
		log.Fatal("fail to add parcel")
	}
	if err := o.AddParcel("medium", "8cm", 2, "nelson", "auckland"); err != nil {
		log.Fatal("fail to add parcel")
	}
	if err := o.AddParcel("large", "13cm", 2, "nelson", "auckland"); err != nil {
		log.Fatal("fail to add parcel")
	}
}
