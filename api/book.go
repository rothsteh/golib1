package api

import (
	"encoding/json"
	"net/http"
)

//Book type
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	//define book
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)

	if err != nil {
		panic(err)
	}
	return ToJSON
}

//FromJson to be used for unmarshalling of book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}

	return book
}

var Books = []Book{
	Book{Title: "Book 1", Author: "joe smith", ISBN: "000111"},
	Book{Title: "Book 2", Author: "New York Mets", ISBN: "000222"},
}

//Bookhandlefunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)

}
