// Filename: cmd/api/quotes.go
package main

import (
	"fmt"
	"net/http"

	// import the data package which contains the definition for Quote
	"github.com/Syha-01/qod/internal/data"
	"github.com/Syha-01/qod/internal/validator"
)

func (a *application) createQuoteHandler(w http.ResponseWriter, r *http.Request) {
	// create a struct to hold a quote
	// we use struct tags[``] to make the names display in lowercase
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}
	// perform the decoding
	// perform the decoding
	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	quote := &data.Quote{
		Content: incomingData.Content,
		Author:  incomingData.Author,
	}

	v := validator.New()

	data.ValidateQuote(v, quote)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors) // implemented later
		return
	}

	// for now display the result
	fmt.Fprintf(w, "%+v\n", incomingData)
}
