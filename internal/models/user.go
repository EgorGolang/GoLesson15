package models

type User struct {
	ID    int    ` json:"id" db:"id"`
	Name  string ` json:"name" db:"name"`
	Email string ` json:"email" db:"email"`
	Age   int    ` json:"age" db:"age"`
}
