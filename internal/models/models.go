package models


type Model struct {
	Id 	  string		`json:"id"`
	CreatedAt string	`json:"created_at"`
	UpdatedAt string	`json:"updated_at"`
	DeletedAt string	`json:"deleted_at"`
}

type User struct {
	Model
	FirstName string	`json:"firstname"`
	LastName  string	`json:"lastname"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
	Pin 	 string		`json:"pin"`
	Wallet    string	`json:"wallet"`

}

type Blacklist struct {
	Email     string
	Token     string
	CreatedAt string
}
