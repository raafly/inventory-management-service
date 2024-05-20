package listing

import "time"

type Register struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,eqfield=Password"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ReqItem struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type ItemUpdate struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Quantity    int    `json:"quantity"`
}

type ItemResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Quantity    int       `json:"quantity"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryNew struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CategoryUpdate struct {
	Id          string `json:"id" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type HistoryResponse struct {
	Id        int       `json:"id"`
	ItemId    int       `json:"itemId"`
	Action    bool      `json:"action"`
	Quantity  int       `json:"quantity"`
	UpdatedAt time.Time `json:"updatedAt"`
}
