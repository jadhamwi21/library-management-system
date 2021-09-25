package main

import (
	"encoding/json"
	"io/ioutil"
)


func GenerateCustomer(Fullname string,Email string,Password string) Customer{
	NewCustomer := Customer{Fullname,Email,Password,BookItems{},0.0}
	return NewCustomer
}
func SaveCustomer(cust* Customer){
	Path := GetPath(cust.Email)
	JSON,_ := json.Marshal(cust)
	ioutil.WriteFile(Path,JSON,0644)
}
func CreateAccountForm() (string, string ,string){
	reader := GetReaderInstance()
	FullName,_ := GetInput("Full Name : ",reader)
	Email,_ := GetInput("Email : ",reader)
	for !ValidateEmail(Email){
		Email,_ = GetInput("Invalid Email Address.\nEmail : ",reader)
	}
	Password := GetPassword()
	return FullName,Email,Password
}
func CreateAccount(){
	Fullname,Email,Password := CreateAccountForm()
	Customer := GenerateCustomer(Fullname,Email,Password)
	SaveCustomer(&Customer)
}