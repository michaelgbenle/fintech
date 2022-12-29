package ports

type Repository interface{
	FindUserByEmail(email string) (*model.Student, error)
	TokenInBlacklist(token *string) bool

}