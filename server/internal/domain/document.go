package domain

import (
	"time"
)

// TransportInvoice - основная модель товарной накладной
type TransportInvoice struct {
	ID        int       `json:"id" db:"id"`
	Number    string    `json:"number" db:"number"`
	Date      time.Time `json:"date" db:"date"`
	Status    string    `json:"status" db:"status"` // draft, completed
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
