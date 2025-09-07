// Filename: internal/data/quotes.go
package data

import (
	"time"

	"github.com/Syha-01/qod/internal/validator"
)

// make our JSON keys be displayed in all lowercase
// "-" means don't show this field

type Quote struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}

func ValidateQuote(v *validator.Validator, quote *Quote) {
	// check if the Content field is empty
	v.Check(quote.Content != "", "content", "must be provided")
	// check if the Author field is empty
	v.Check(quote.Author != "", "author", "must be provided")
	// check if the Content field is empty
	v.Check(len(quote.Content) <= 100, "content", "must not be more than 100 bytes long")
	// check if the Author field is empty
	v.Check(len(quote.Author) <= 25, "author", "must not be more than 25 bytes long")
}

// BODY='{"content":"I loved it!💛", "author":"Dalwin D. Lewis"}'
// curl -i -d "$BODY" localhost:4000/v1/quotes
