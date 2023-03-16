package votes

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents a vote builder
type Builder interface {
	Create() Builder
	IsAccepted() Builder
	IsDenied() Builder
	IsNeutral() Builder
	Now() (Vote, error)
}

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	IsAccepted() bool
	IsDenied() bool
	IsNeutral() bool
}
