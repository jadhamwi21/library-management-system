package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func BrowseBooks() {
	var Books BookItems
	Content,_ := os.ReadFile( GetBooksJsonPath())
	json.Unmarshal(Content,&Books)
	Flags := BrowseEnums{true,true,true}
	FormattedString := Books.FormatBooks(Flags)
	fmt.Println(FormattedString)
}