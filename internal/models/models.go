package models

type User struct {
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
