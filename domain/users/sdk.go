package users

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

// Builder represents a user builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithProfile(profile string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (User, error)
}

// User represents a user
type User interface {
	Hash() hash.Hash
	Name() string
	PubKey() hash.Hash
	Profile() string
	CreatedOn() time.Time
}
