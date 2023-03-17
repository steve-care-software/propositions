package users

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a user builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPubKey(pubKey hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (User, error)
}

// User represents a user
type User interface {
	Hash() hash.Hash
	Name() string
	PubKey() hash.Hash
	CreatedOn() time.Time
}
