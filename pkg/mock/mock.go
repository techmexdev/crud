package mock

import (
	"errors"

	"github.com/techmexdev/crud"
)

// NewUserStore returns a new UserStorage.
func NewUserStore() crud.UserStorage {
	return &UserStorage{users: []crud.User{}}
}

// UserStorage is a mock implementation
// of crud.UserStorage.
type UserStorage struct {
	users []crud.User
}

// Create inserts a user in the storage.
func (us *UserStorage) Create(crud.User) error {
	return errors.New("not implemented yet")
}

// Read retrieves a user from the storage.
func (us *UserStorage) Read(username string) (crud.User, error) {
	return crud.User{}, errors.New("not implemented yet")
}

// Update changes a user's properties.
func (us *UserStorage) Update(username string, u crud.User) error {
	return errors.New("not implemented yet")
}

// Delete removes a user from the storage.
func (us *UserStorage) Delete(username string) error {
	return errors.New("not implemented yet")
}

// List retrieves all users from the storage.
func (us *UserStorage) List() ([]crud.User, error) {
	return []crud.User{}, errors.New("not implemented yet")
}
