package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("ProcessTruck", func(t *testing.T) {
		t.Run("Should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1", cargo: 42}
			et := &ElectricTruck{id: "2"}
			if err := processTruck(nt); err != nil {
				t.Fatalf("error  processing truck: %s", err)
			}

			if err := processTruck(et); err != nil {
				t.Fatalf("error  processing truck: %s", err)
			}

			fmt.Println(nt.cargo)
			if nt.cargo != 0 {
				t.Fatal("Normal truck cargo should be 0")
			}

		})
	})
}
