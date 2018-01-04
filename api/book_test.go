package api

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "Howie Rothstein", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"Howie Rothstein","isbn":"0123456789"}`,
		string(json), "Book JSON marshalling wrong.")
}
func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"Title":"Cloud Native Go","Author":"Howie Rothstein","ISBN":"0123456789"}`)
	fmt.Println("json is:", string(json))
	fmt.Println("From JSON is:", FromJSON(json))
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "Howie Rothstein", ISBN: "0123456789"}, book, "Book unmarshalling wrong")
}
