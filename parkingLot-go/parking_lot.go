package main

import "errors"

type VehicleType = int

const (
	CAR VehicleType = iota
	MOTO
)

type Vehicle struct {
	Type VehicleType
}

type Spot struct {
	SpotType VehicleType
	Ocupied  bool
	Vehicle  Vehicle
}

func (s *Spot) NewSpot(vehicleType VehicleType) *Spot {
	return &Spot{
		SpotType: vehicleType,
		Ocupied:  false,
		Vehicle:  Vehicle{},
	}
}

func (s *Spot) Park(vehicle Vehicle) error {
	if s.Ocupied {
		return errors.New("Cant park. Slot was ocupied")
	} else if vehicle.Type == CAR && s.SpotType == MOTO {
		return errors.New("Cant park. Can't fit car into a Moto space")
	}
	s.Vehicle = vehicle
	s.Ocupied = true
	return nil

}
func (s *Spot) UnPark() error {
	if !s.Ocupied {
		return errors.New("Free spot. Can't unpark")
	}
	s.Vehicle = Vehicle{}
	s.Ocupied = false
	return nil

}

type ParkingLot struct {
	Levels [][][]*Spot
}

func NewParkingLot(levels [][][]*Spot) *ParkingLot {
    return &ParkingLot{Levels: levels}
}
func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) error {
	for _, level := range pl.Levels {
		for _, row := range level {
			for _, spot := range row {
                if err := spot.Park(vehicle); err == nil {
                    return nil 
                }
			}
		}
	}
	return errors.New("no available parking slots")
}

func (pl *ParkingLot) UnparkVehicle(vehicle *Vehicle) error {
    for _, level := range pl.Levels {
        for _, row := range level {
            for _, spot := range row {
                if err := spot.UnPark(); err == nil {
                    return nil 
                }
            }
        }
    }
	return errors.New("This vehicle was not parked at this parking lot")
}
func (pl *ParkingLot) FindVehicleBySpot(spot *Spot) Vehicle {
    return spot.Vehicle
}

func main() {

}
