package crud

// User is a crud app user.
type User struct {
	FirstName string
	LastName  string
	Username  string
}

// UserStorage stores and retrieves users.
type UserStorage interface {
	Create(User) error
	Read(username string) (User, error)
	Update(username string, u User) (err error)
	Delete(username string) error
	List() ([]User, error)
}
