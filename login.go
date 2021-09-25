package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetCustomerFile(Email string) ([]byte , error){
	File,err := os.ReadFile(GetPath(Email))
	return File,err
}
func PasswordProcess(Password* string, RealCustomerPassword string){
	reader := GetReaderInstance()
	for *Password != RealCustomerPassword{
		fmt.Println("Password Doesn't Match")
		*Password,_ = GetInput("Password : ",reader)
	}
}
func ParseCustomerJsonData(JSON []byte) Customer{
	var Cust Customer
	json.Unmarshal(JSON,&Cust)
	return Cust
}


func LoginForm() (string,string){
	reader := GetReaderInstance()
ReFill:
	Email, _ := GetInput("Email : ", reader)
	Password, _ := GetInput("Password : ", reader)
	if CheckIfAdmin(Email,Password){
		return Email,Password
	}
	if !ValidateEmail(Email) {
		fmt.Println("Invalid Email Format")
		goto ReFill
	}
	return Email,Password
}
func CheckIfAdmin(Email string,Password string)bool{
	return (Email == Password && Email == "admin")
}
func Login() (*Customer,bool){
Retry:
	Email,Password := LoginForm()
	if CheckIfAdmin(Email,Password){
		return nil,true
	}
	Content,err := GetCustomerFile(Email)
	if err != nil {
		fmt.Println("User Doesn't Exist")
		goto Retry
	}
	 customer := ParseCustomerJsonData(Content) 
	 PasswordProcess(&Password,customer.Password)
	 return &customer,false
}