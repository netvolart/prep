import { describe, it, expect, beforeEach } from "@jest/globals";
import { ParkingLot, Spot, Vehicle, VehicleType } from ".";

describe("ParkingLot", () => {
  let lot: ParkingLot;
  let car: Vehicle;
  let moto: Vehicle;

  beforeEach(() => {
    const levels = [
      [
        [
          new Spot(VehicleType.MOTO),
          new Spot(VehicleType.MOTO),
          new Spot(VehicleType.CAR),
        ],
        [
          new Spot(VehicleType.CAR),
          new Spot(VehicleType.CAR),
          new Spot(VehicleType.MOTO),
        ],
      ],
      [
        [
          new Spot(VehicleType.CAR),
          new Spot(VehicleType.CAR),
          new Spot(VehicleType.MOTO),
        ],
      ],
    ];
    lot = new ParkingLot(levels);
    car = new Vehicle(VehicleType.CAR);
    moto = new Vehicle(VehicleType.MOTO);
  });

  it("should park a car in a car spot", () => {
    expect(lot.parkVehicle(car)).toBe(true);
  });

  it("should park a moto in any spot", () => {
    expect(lot.parkVehicle(moto)).toBe(true);
  });

  it("should not park a car in a moto spot", () => {
    const spot = new Spot(VehicleType.MOTO);
    expect(spot.park(car)).toBe(false);
  });

  it("should unpark a vehicle", () => {
    lot.parkVehicle(car);
    expect(lot.unparkVehicle(car)).toBe(true);
  });

  it("should return null if spot is empty", () => {
    const spot = new Spot(VehicleType.CAR);
    expect(spot.unpark()).toBeNull();
  });

  it("should find vehicle by spot", () => {
    lot.parkVehicle(car);

    let foundSpot: Spot | null = null;
    for (const level of lot.levels) {
      for (const row of level) {
        for (const spot of row) {
          console.log(spot);
          if (spot.vehicle === car) {
            console.log(spot);
            foundSpot = spot;
          }
        }
      }
    }
    expect(foundSpot).not.toBeNull();
    expect(lot.findVehiculeBySpot(foundSpot!)).toBe(car);
  });
});
