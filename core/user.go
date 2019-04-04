package core

type (
	User struct {
		ID        int64
		UpdatedAt string
		CreatedAt string
		Email     string
		Password  string
	}
	UserStore interface {
		Find(int64) (*User, error)
		Create(user *User) error
	}

	UserService interface {
		Register(user *User)
	}
)
