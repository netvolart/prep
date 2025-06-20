export enum VehicleType {
  CAR,
  MOTO,
}

export class Vehicle {
  constructor(public type: VehicleType) {}
}

export class Spot {
  public occupied: boolean = false;
  public vehicle: Vehicle | null = null;

  constructor(public spotType: VehicleType) {}

  public park(vehicle: Vehicle): boolean {
    if (this.occupied) {
      console.log("Slot is ocupied");
      return false;
    } else if (
      vehicle.type === VehicleType.CAR &&
      this.spotType === VehicleType.MOTO
    ) {
      console.log("Unable to park Car into a Moto spot");
      return false;
    }
    this.vehicle = vehicle;
    this.occupied = true;
    console.log(`Parked ${VehicleType[vehicle.type]} into a spot`);
    return true;
  }

  public unpark(): Vehicle | null {
    if (!this.occupied) {
      console.log("This Spot is not ocupied by vehicule.");
      return null;
    }
    this.vehicle = null;
    this.occupied = false;
    return this.vehicle;
  }
}

export class ParkingLot {
  constructor(public levels: Spot[][][]) {}

  public parkVehicle(vehicle: Vehicle): boolean {
    for (const level of this.levels) {
      for (const row of level) {
        for (const spot of row) {
          if (spot.park(vehicle)) return true;
        }
      }
    }
    console.log("No avaliable parking slots");
    return false;
  }

  public unparkVehicle(vehicle: Vehicle): boolean {
    for (const level of this.levels) {
      for (const row of level) {
        for (const spot of row) {
          if (spot.vehicle === vehicle) {
            spot.unpark()
            return true;
          }
        }
      }
    }
    console.log('This vehicule was not parket at this parkingLot')
    return false
  }


  public findVehiculeBySpot(spot: Spot): Vehicle | null {
    return spot.vehicle;
  }
}