package main

import (
	"fmt"
	"strconv"
)

func ChargeAccount(cust *Customer) *Customer{
	reader := GetReaderInstance()
Back:
	Input, _ := GetInput("How Much You Want To Charge Your Account ?\n", reader)
	parsedAmount, err := strconv.ParseFloat(Input,64)
	if err == nil{
		cust.Balance += parsedAmount
		UpdateCustomerFile(cust)
		return cust
	}else{
		fmt.Println("Amount Should Be A Number")
		goto Back
	}
}