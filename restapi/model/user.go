package model

type User struct {
	Email     string `json:"email" binding:"required"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
}

type ErrorMessage struct {
	ErrCode string `json:"errCode"`
	Message string `json:"message" binding:"required"`
}
