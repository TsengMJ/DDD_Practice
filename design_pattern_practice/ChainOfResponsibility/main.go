package main

import "chain_of_responsibility/department"

func main() {

	cashier := &department.Cashier{}

	//Set next for medical department
	medical := &department.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &department.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &department.Reception{}
	reception.SetNext(doctor)

	patient := &department.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
