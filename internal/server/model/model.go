package model

type Status string

const (
	StatusActive Status = "active"
	StatusInactive Status = "inactive"
	StatusPending Status = "pending"
)

type User struct {
	ID string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Phone string `json:"phone" db:"phone"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Password string `json:"password" db:"password"`
	Status Status `json:"status" db:"status"`
	LastLoginAt string `json:"last_login_at" db:"last_login_at"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}