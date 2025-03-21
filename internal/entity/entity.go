package entity

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
}
