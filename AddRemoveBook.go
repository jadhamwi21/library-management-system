package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)
func GetBooksJsonPath() string{
	return "./Data/books.json"
}
func GetBooksTree() ([]byte,error){
	Content,err := os.ReadFile(GetBooksJsonPath())
	return Content,err
}
func CheckBookIDExistance(bookId int,Books BookItems)bool{
	_,ok := Books[bookId]
	return ok
}
func AddBookToFile(book Book){
	BooksPath := GetBooksJsonPath()
	JsonTree,err := GetBooksTree()
	if err != nil{
		books := BookItems{book.BookId : book}
		Json,_ := json.Marshal(books)
		os.WriteFile(BooksPath,Json,0644)
	}else{
		var books BookItems
		json.Unmarshal(JsonTree,&books)
		for CheckBookIDExistance(book.BookId,books){
			reader := GetReaderInstance()
			fmt.Println("ID Already Exists Enter Another ID")
			IdInput,_ := GetInput("Book ID : ",reader)
			ParsedID,err := strconv.Atoi(IdInput)
			if err == nil{
				book.BookId = ParsedID
			}
		}
		books[book.BookId] = book
		newJsonTree,_ := json.Marshal(books)
		os.WriteFile(BooksPath,newJsonTree,0644)
		fmt.Println("Book Added Successfully")
	}
}
func BookForm() (int,string,string){
	reader := GetReaderInstance()
Retry:
	ID,_ := GetInput("Book ID : ",reader)
	idParsed,parseErr := strconv.Atoi(ID)
	if parseErr != nil{
			fmt.Println("ID Must Be An Integer")
			goto Retry
		}
	Title,_ := GetInput("Book Title : ",reader)
	Author,_ := GetInput("Book Author : ",reader)

	return idParsed,Title,Author
}

func CreateBookType(id int,author string,title string) Book{
	book := Book{BookId: id,BookAuthor:author,BookTitle: title}
	return book
}

func AddBook(){
	ID,Title,Author :=BookForm()
	Book := CreateBookType(ID,Title,Author)
	AddBookToFile(Book)
}
func RemoveBook(){
	reader := GetReaderInstance()
Retry:
	BookID,_ := GetInput("Enter Book ID Please",reader)
	ParsedBookID,err := strconv.Atoi(BookID)
	if err != nil{
		fmt.Println("ID Must Be An Integer")
		goto Retry
	}
	BooksTree,_ := GetBooksTree()
	var parsedBooksTree BookItems
	json.Unmarshal(BooksTree,&parsedBooksTree)
	if _,ok := parsedBooksTree[ParsedBookID] ; ok{
		delete(parsedBooksTree,ParsedBookID)
		fmt.Println("Book Removed Successfully")
	}else{
		fmt.Println("Book Doesn't Exist In The Database");
		return
	}
	newJsonTree,_ := json.Marshal(parsedBooksTree)
	os.WriteFile(GetBooksJsonPath(),newJsonTree,0644)
}


func AddRemoveBook(){
	reader := GetReaderInstance()
Retry:
	Message := "1 - Add\n2 - Remove\n"
	Input,_ := GetInput(Message,reader)
	switch Input{
	case "1":
		AddBook()
	case "2":
		RemoveBook()
	default:
		fmt.Println("Invalid Selection")
		goto Retry
	}
}