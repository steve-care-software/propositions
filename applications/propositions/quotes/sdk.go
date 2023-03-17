package quotes

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/applications/propositions/quotes/managers"
	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/quotes"
)

// Builder represents a quote application builder
type Builder interface {
	Create() Builder
	WithProposition(proposition propositions.Proposition) Builder
	Now() (Application, error)
}

// Application represents a quote application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (quotes.Quote, error)
	Insert(quote quotes.Quote) error
	Update(original hash.Hash, updated quotes.Quote) error
	Delete(hash hash.Hash) error
	Manager(quote quotes.Quote) (managers.Application, error)
}
