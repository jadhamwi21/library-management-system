package main

import (
	"fmt"
)

func UserEnterInterface() (*Customer,bool){
	reader := GetReaderInstance()
	fmt.Println("Welcome To Library Management System")
	var InterfaceMessage string = "1 - Create Account\n2 - Login\n"
Interface:
	Input,_ := GetInput(InterfaceMessage,reader)
	switch Input {
	case "1":
		CreateAccount()
		goto Interface
	case "2":
		return Login()

	default:
		fmt.Println("Invalid Selection, Retry")
		goto Interface
	}
}
func AdminInterface(){
	reader := GetReaderInstance()
Back:
	InterfaceMessage := "1 - Add / Remove Book\n2 - View Customer Informations\n"
	Input,_ := GetInput(InterfaceMessage,reader)
	switch Input{
	case "1":
		AddRemoveBook()
		goto Back
	case "2":
		ViewCustomerInformations()
		goto Back
	default:
		fmt.Println("Invalid Option")
		goto Back
	}
}
func CustomerInterface(cust* Customer){
	reader := GetReaderInstance()
Back:
	InterfaceMessage := "1 - Browse Books\n2 - Borrow Book\n3 - Charge Account\n4 - View Account Info\n"
	Input,_ := GetInput(InterfaceMessage,reader)
	switch Input{
	case "1":
		BrowseBooks()
		goto Back
	case "2":
		cust = BorrowBook(cust)
		goto Back
	case "3":
		cust = ChargeAccount(cust)
		goto Back
	case "4":
		fmt.Println(cust.Format())
		goto Back

	default:
		fmt.Println("Invalid Option")
		goto Back
	}
}