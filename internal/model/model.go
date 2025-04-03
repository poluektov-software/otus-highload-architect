package model

type Account struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//CreatedAt string `json:"created_at"`
	//UpdatedAt string `json:"updated_at"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
