package main

import (
	"errors"
	"fmt"
	"log"
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

	if err := truck.LoadCarGo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	return nil
}

func main() {

	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "2"}
	if err := processTruck(nt); err != nil {
		log.Fatalf("error  processing truck: %s", err)
	}

	if err := processTruck(et); err != nil {
		log.Fatalf("error  processing truck: %s", err)
	}

	log.Println(nt.cargo)
	log.Println(et.battery)
}
