package managers

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/applications/propositions/quotes/managers/tasks"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
)

// Builder represents a manager application builder
type Builder interface {
	Create() Builder
	WithQuote(quote quotes.Quote) Builder
	Now() (Application, error)
}

// Application represents a manager application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (managers.Manager, error)
	Insert(manager managers.Manager) error
	Update(original hash.Hash, updated managers.Manager) error
	Delete(hash hash.Hash) error
	Task(managers managers.Manager) (tasks.Application, error)
}
