package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("Truck not found")
)

func (t *Truck) LoadCarGo() error {
	return ErrTruckNotFound
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCarGo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	return ErrNotImplemented
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		err := processTruck(truck)
		if err != nil {
			if errors.Is(err, ErrTruckNotFound) {
				log.Fatal("True")
			}

			// Fatalf is equivalent to [Printf] followed by a call to [os.Exit](1).
			log.Fatalf("error processing truck: %s", err)

		}

		//use this syntax so var err is local not global
		// if err := processTruck(truck); err != nil {
		// 	log.Fatalf("error processing truck: %s", err)
		// }
	}
}
