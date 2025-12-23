package domain

// Signer represents a signer entity.
type Signer struct {
	ID         int    `json:"id" db:"id"`
	DocumentID int    `json:"document_id" db:"document_id"`
	Name       string `json:"name" db:"name"`
	Position   string `json:"position" db:"position"`
	Email      string `json:"email" db:"email"`
	Phone      string `json:"phone" db:"phone"`
}
