package propositions

import (
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/applications/propositions/quotes"
	propositions "github.com/steve-care-software/propositions/domain"
)

// Builder represents a proposition builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	Now() (Application, error)
}

// Application represents a proposition application
type Application interface {
	List() []hash.Hash
	Retrieve(hash hash.Hash) (propositions.Proposition, error)
	Update(original hash.Hash, updated propositions.Proposition) error
	Delete(hash hash.Hash) error
	Quote(proposition propositions.Proposition) (quotes.Application, error)
}
