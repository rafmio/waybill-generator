package domain

// Vehicle представляет единое транспортное средство (тягач + прицеп)
type Vehicle struct {
	ID int `json:"id"` // Уникальный идентификатор записи

	// Данные тягача
	TractorBrand      string `json:"tractor_brand"`      // Марка тягача
	TractorModel     string `json:"tractor_model"`     // Модель тягача
	TractorYear      int    `json:"tractor_year"`      // Год выпуска тягача
	TractorLicensePlate string `json:"tractor_license_plate"` // Гос. номер тягача

	// Данные прицепа
	TrailerBrand      string `json:"trailer_brand,omitempty"`      // Марка прицепа (может отсутствовать)
	TrailerModel     string `json:"trailer_model,omitempty"`     // Модель прицепа (может отсутствовать)
	TrailerYear      int    `json:"trailer_year,omitempty"`      // Год выпуска прицепа (может отсутствовать)
	TrailerLicensePlate string `json:"trailer_license_plate,omitempty"` // Гос. номер прицепа (может отсутствовать)

	// Общие метаданные
	RegistrationDate *time.Time `json:"registration_date,omitempty"` // Дата регистрации (если нужна)
	Notes          string      `json:"notes,omitempty"`               // Дополнительные примечания
}
