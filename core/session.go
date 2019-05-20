package core

type (
	SessionService interface {
		FindUserByEmail(email string) (*User, error)
	}
)
