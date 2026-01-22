package domain

import "time"

// Amiibo represents a Nintendo Amiibo figure.
type Amiibo struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ImageURL    string    `json:"image_url"`
	Series      string    `json:"series"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
}
