package main

import (
    "testing"
)

func TestParkVehicleSuccess(t *testing.T) {
    // Create a parking lot with one car spot and one moto spot
    levels := [][][]*Spot{
        {
            {&Spot{SpotType: CAR}, &Spot{SpotType: MOTO}},
        },
    }
    lot := NewParkingLot(levels)
    car := Vehicle{Type: CAR}
    moto := Vehicle{Type: MOTO}

    if err := lot.ParkVehicle(car); err != nil {
        t.Errorf("expected to park car, got error: %v", err)
    }
    if !levels[0][0][0].Ocupied {
        t.Errorf("expected car spot to be occupied")
    }

    if err := lot.ParkVehicle(moto); err != nil {
        t.Errorf("expected to park moto, got error: %v", err)
    }
    if !levels[0][0][1].Ocupied {
        t.Errorf("expected moto spot to be occupied")
    }
}

func TestParkVehicleFailWhenFull(t *testing.T) {
    levels := [][][]*Spot{
        {
            {&Spot{SpotType: CAR, Ocupied: true}, &Spot{SpotType: MOTO, Ocupied: true}},
        },
    }
    lot := NewParkingLot(levels)
    car := Vehicle{Type: CAR}
    if err := lot.ParkVehicle(car); err == nil {
        t.Errorf("expected error when parking in full lot, got nil")
    }
}

func TestParkVehicleCarInMotoSpot(t *testing.T) {
    levels := [][][]*Spot{
        {
            {&Spot{SpotType: MOTO}},
        },
    }
    lot := NewParkingLot(levels)
    car := Vehicle{Type: CAR}
    if err := lot.ParkVehicle(car); err == nil {
        t.Errorf("expected error when parking car in moto spot, got nil")
    }
}

func TestUnparkVehicle(t *testing.T) {
    spot := &Spot{SpotType: CAR}
    levels := [][][]*Spot{{{spot}}}
    lot := NewParkingLot(levels)
    car := Vehicle{Type: CAR}
    _ = lot.ParkVehicle(car)

    if err := lot.UnparkVehicle(&car); err != nil {
        t.Errorf("expected to unpark car, got error: %v", err)
    }
    if spot.Ocupied {
        t.Errorf("expected spot to be unoccupied after unparking")
    }
}

func TestUnparkVehicleNotFound(t *testing.T) {
    levels := [][][]*Spot{
        {
            {&Spot{SpotType: CAR}},
        },
    }
    lot := NewParkingLot(levels)
    car := Vehicle{Type: CAR}
    if err := lot.UnparkVehicle(&car); err == nil {
        t.Errorf("expected error when unparking vehicle not in lot, got nil")
    }
}

func TestFindVehicleBySpot(t *testing.T) {
    spot := &Spot{SpotType: CAR}
    levels := [][][]*Spot{{{spot}}}
    lot := NewParkingLot(levels)
    car := Vehicle{Type: CAR}
    _ = lot.ParkVehicle(car)

    found := lot.FindVehicleBySpot(spot)
    if found.Type != CAR {
        t.Errorf("expected to find car in spot, got %v", found.Type)
    }
}