package domain

// GoodsItem - товарная позиция
type GoodsItem struct {
	ID          int     `json:"id" db:"id"`
	DocumentID  int     `json:"document_id" db:"document_id"`
	Name        string  `json:"name" db:"name"`
	Quantity    float64 `json:"quantity" db:"quantity"` // значение в единицах измерения
	Unit        string  `json:"unit" db:"unit"`         // единица измерения: кг, шт, м3
	Description string  `json:"description" db:"description"`
}
