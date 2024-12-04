package filter

import "time"

type GroupFilter struct {
	Name       string     `json:"name" filter:"1" searchable:"1"`
	CreatedAt  *time.Time `json:"created_at" filter:"2"` // Filtro per la data di creazione
	UpdatedAt  *time.Time `json:"updated_at" filter:"2"` // Filtro per la data di aggiornamento
	Permission string     `json:"permission" filter:"1" field_filter:"name"`
	SortBy     string     `json:"sort_by" filter:"4"`    // Campo su cui ordinare
	SortOrder  string     `json:"sort_order" filter:"5"` // Ordine di ordinamento (asc/desc)
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	Search     string     `json:"search"`
}
