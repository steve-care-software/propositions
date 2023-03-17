package tasks

import (
	"time"

	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a task builder
type Builder interface {
	Create() Builder
	WithQuote(quote quotes.Quote) Builder
	WithManagers(managers managers.Managers) Builder
	WithDeposit(deposit units.Unit) Builder
	CreatedOn(createdOn time.Time) Builder
	ExpiresOn(expiresOn time.Time) Builder
	Now() (Task, error)
}

// Task represents a task
type Task interface {
	Hash() hash.Hash
	Quote() quotes.Quote
	Managers() managers.Managers
	Deposit() units.Unit
	CreatedOn() time.Time
	ExpiresOn() time.Time
}
