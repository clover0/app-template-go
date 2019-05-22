package core

type (
	SessionService interface {
		CreateSession(userId uint32) (string, error)
		GetSession(string) (string, error)
		ComparePassword(user *User, input string) error
		FindUserByEmail(email string) (*User, error)
		FindUserById(id uint32) (*User, error)
	}
)
