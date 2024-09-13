package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	_, err := BuildFactory(3)
	if err == nil {
		t.Error("An error must be returned for an invalid factory type")
	}

	motorbikeF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	motorbikeVehicle, err := motorbikeF.GetVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n",
		motorbikeVehicle.NumWheels())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.NumSeats())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.NumWheels())

	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())
	motorbikeVehicle, err = motorbikeF.GetVehicle(CruiseMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Motorbike vehicle has %d wheels\n",
		motorbikeVehicle.NumWheels())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.NumSeats())
	t.Logf("Motorbike vehicle has %d seats\n", motorbikeVehicle.NumWheels())

	cruiseBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", cruiseBike.GetMotorbikeType())
	_, err = motorbikeF.GetVehicle(3)
	if err == nil {
		t.Error("An error must be returned for an invalid vehicle type")
	}

}

func TestCarFactory(t *testing.T) {
	carF, err := BuildFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.GetVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
	t.Logf("Luxury car has %d seats.\n", luxuryCar.(Vehicle).NumSeats())
	carVehicle, err = carF.GetVehicle(FamilyCarType)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Car vehicle has %d seats\n", carVehicle.NumWheels())
	familyCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Family car has %d doors.\n", familyCar.NumDoors())
	t.Logf("Family car has %d seats.\n", familyCar.(Vehicle).NumSeats())

	_, err = carF.GetVehicle(3)
	if err == nil {
		t.Fatal("GetVehicle must return an error")
	}

}
