package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func SearchBook(bookId int) (Book,error){
	Bookspath := GetBooksJsonPath()
	var Books BookItems
	 Content,_ := os.ReadFile(Bookspath)
	json.Unmarshal(Content,&Books)
	return Books.FindBookById(bookId)
}

func BorrowBook(cust *Customer) *Customer{
	reader := GetReaderInstance()
	BrowseBooks()
Back:
	Input, _ := GetInput("Enter Book ID : ", reader)
	parsedId,err := strconv.Atoi(Input)
	if err != nil{
		fmt.Println("Book ID Must Be An Integer")
		goto Back
	}
	Book,err := SearchBook(parsedId)
	if err == nil{
		cust.BorrowedBooks.AddBookToList(Book)
	}else{
		fmt.Println("This Book Doesn't Exist Please Enter A Valid Book ID")
		goto Back
	}
	UpdateCustomerFile(cust)
	return cust
}