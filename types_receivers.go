package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)
type BrowseEnums struct{
	ID bool
	Title bool
	Author bool
}

type Book struct {
	BookId     int
	BookAuthor string
	BookTitle  string
}
type BookItems map[int]Book

type Customer struct {
	Fullname      string
	Email         string
	Password      string
	BorrowedBooks BookItems
	Balance       float64
}
func (customer Customer) Format() string{
	formattedString := ""
	Type := reflect.TypeOf(customer)
	Value := reflect.ValueOf(customer)
	for i := 0; i < Value.NumField(); i++ {
		Key := Type.Field(i).Name
		Value := Value.Field(i)
		if(Key == "Password"){
			continue
		}
		if(Key == "BorrowedBooks"){
			Flags := BrowseEnums{
				Author: true,
				ID: false,
				Title:true,
			}
			formattedString += Value.Interface().(BookItems).FormatBooks(Flags)
			continue
		}
		formattedString += fmt.Sprintf("%-10v%-10v\n", Key+" : ", Value)
	}
	return strings.TrimSuffix(formattedString,"\n")
}

func (bookItems BookItems) FormatBooks(enum BrowseEnums)string{
	BooksIterator := reflect.ValueOf(bookItems).MapRange()
	NumberOfFields := BoolToInt(enum.Author) + BoolToInt(enum.ID) + BoolToInt(enum.Title)
	Args := []interface{}{
	}
	if(enum.ID){
		Args = append(Args, "Book ID")
	}
	if(enum.Title){
		Args = append(Args, "Book Title")
	}
	if(enum.Author){
		Args = append(Args,"Book Author")
	}
	FormattedString := fmt.Sprintf(strings.Repeat("%-30v",NumberOfFields),Args...)
	FormattedString += "\n"
	for BooksIterator.Next(){
		if(enum.ID){
			FormattedString += fmt.Sprintf("%-30v",BooksIterator.Value().Field(0))
		}
		if(enum.Author){
			FormattedString += fmt.Sprintf("%-30v",BooksIterator.Value().Field(1))
		}
		if(enum.Title){
			FormattedString += fmt.Sprintf("%-30v",BooksIterator.Value().Field(2))
		}
		FormattedString += "\n"
	}
	return FormattedString
}
func (bookItems BookItems) FindBookById(id int) (Book,error){
	MapIterator := reflect.ValueOf(bookItems).MapRange()
	for MapIterator.Next(){
		Value := MapIterator.Value().Interface().(Book)
		if(Value.BookId == id){
			return Value,nil
		}
	}
	return Book{},errors.New("couldn't find a book with this id")
}
func (bookItems* BookItems) AddBookToList(book Book){
	books := *bookItems
	books[book.BookId] = book
}