package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"reflect"
	"strings"

	"github.com/howeyc/gopass"
)
func BoolToInt(value bool) int{
	if(value){
		return 1
	}else{
		return 0
	}
}
func UpdateCustomerFile(cust* Customer){
	Path := GetPath(cust.Email)
	JSON,_ := json.Marshal(cust)
	os.WriteFile(Path,JSON,0644)
}

func GetReaderInstance() *bufio.Reader{
	reader := bufio.NewReader(os.Stdin)
	return reader
}
func GetPath(CustomerEmail string) string{
	return fmt.Sprintf("./Data/%v.json",CustomerEmail)
}
func FormatBookStruct(bookStruct reflect.Value)string{
	return fmt.Sprintf("%-30v%-30v",bookStruct.Field(1).String(),bookStruct.Field(2).String())
}

func GetInput(prompt string, Reader *bufio.Reader) (string,error){
	fmt.Print(prompt)
	input,error := Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input,error
}
func ValidateEmail(email string) bool{
	_,err := mail.ParseAddress(email)
	return err == nil
}
func GetPassword() string{
	fmt.Print("Password : ")
	password,_ := gopass.GetPasswd()
	return string(password)
}
