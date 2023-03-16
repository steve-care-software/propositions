package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/users"
)

// Builder represents a proposition builder
type Builder interface {
	Create() Builder
	WithPubKey(pubKey hash.Hash) Builder
	WithTitle(title string) Builder
	WithDescription(description string) Builder
	WithParent(parent Proposition) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Proposition, error)
}

// Proposition represents a proposition
type Proposition interface {
	Hash() hash.Hash
	SubmittedBy() users.User
	Title() string
	Description() string
	CreatedOn() time.Time
	HasParent() bool
	Parent() Proposition
}
