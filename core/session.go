package core

type (
	SessionService interface {
		FindUserByEmail(email string) (*User, error)
		FindUserById(id uint32) (*User, error)
	}
)
