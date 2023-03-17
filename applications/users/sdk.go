package users

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/users"
)

// Application represents a user application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (users.User, error)
}
