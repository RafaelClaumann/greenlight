package data

import (
	"time"

	"greenlight.claumann.net/internal/data"
)

// Annotate the Movie struct with struct tags to control how the keys appear in the JSON-encoded output.
type Movie struct {
	ID        int64        `json:"id"`
	CreatedAt time.Time    `json:"-"` // Use the - directive
	Title     string       `json:"title"`
	Year      int32        `json:"year,omitempty"` // Add the omitempty directive
	Runtime   data.Runtime `json:"runtime,omitempty"`
	Genres    []string     `json:"genres,omitempty"` // Add the omitempty directive
	Version   int32        `json:"version"`
}
