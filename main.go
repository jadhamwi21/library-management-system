package main

func main() {
	Customer, isAdmin := UserEnterInterface()
	if isAdmin {
		AdminInterface()
	} else {
		CustomerInterface(Customer)
	}
}