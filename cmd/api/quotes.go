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

	// Add the comment to the database table
	err = a.quoteModel.Insert(quote)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// Set a Location header. The path to the newly created quote
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/quotes/%d", quote.ID))

	// Send a JSON response with 201 (new resource created) status code
	data := envelope{
		"quote": quote,
	}
	err = a.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

}

//PAGE 186
