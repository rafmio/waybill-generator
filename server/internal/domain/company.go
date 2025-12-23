package domain

type Company struct {
	ID      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	TaxID   string `json:"tax_id" db:"tax_id"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone"`
	Email   string `json:"email" db:"email"`
}
