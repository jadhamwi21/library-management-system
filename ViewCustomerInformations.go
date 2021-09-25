package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
func GetCustomersJSON() []string{
	files,_ := ioutil.ReadDir("./Data")
	CustomersJsonSlice := []string{}
	for _,file := range files{
		if(file.Name() != "books.json"){
		CustomersJsonSlice = append(CustomersJsonSlice,file.Name())}
	}
	return CustomersJsonSlice
}
func SearchCustomer(CustomerFullName string)*Customer{
	CustomersJsonFileNames := GetCustomersJSON()
	for _,CustomerFileName := range CustomersJsonFileNames{
		Content,_ := os.ReadFile(fmt.Sprintf("./Data/%v",CustomerFileName))
		var cust Customer
		json.Unmarshal(Content,&cust)
		if(cust.Fullname == CustomerFullName){
			return &cust
		}
	}
	return nil
}


func ViewCustomerInformations() {
	reader := GetReaderInstance()
	Fullname,_ := GetInput("Full Name : ",reader)
	customer := SearchCustomer(Fullname)
	formattedString := customer.Format()
	fmt.Println(formattedString)
}