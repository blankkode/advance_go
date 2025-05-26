package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Truck interface {
	LoadCarGo() error
	UnLoadCarGo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("Truck not found")
)

func (t *NormalTruck) LoadCarGo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnLoadCarGo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCarGo() error {
	e.cargo += 1
	e.battery += -1
	return nil
}

func (e *ElectricTruck) UnLoadCarGo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

func processTruck(truck Truck) error {

	// %+v will show both property name and value
	fmt.Printf("process truck %+v \n", truck)

	// simulate some processing time
	time.Sleep(time.Second)

	if err := truck.LoadCarGo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnLoadCarGo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed successfully")
}
