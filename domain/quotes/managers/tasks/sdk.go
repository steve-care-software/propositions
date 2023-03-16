package tasks

import (
	"time"

	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
)

// Builder represents a task builder
type Builder interface {
	Create() Builder
	WithQuote(quote quotes.Quote) Builder
	WithManagers(managers managers.Managers) Builder
	WithDeposit(deposit units.Units) Builder
	CreatedOn(createdOn time.Time) Builder
	ExpiresOn(expiresOn time.Time) Builder
	Now() (Task, error)
}

// Task represents a task
type Task interface {
	Hash() hash.Hash
	Quote() quotes.Quote
	Managers() managers.Managers
	Deposit() units.Units
	CreatedOn() time.Time
	ExpiresOn() time.Time
}

// CompletedTaskBuilder represents a completed task builder
type CompletedTaskBuilder interface {
	Create() CompletedTaskBuilder
	WithTask(task Task) CompletedTaskBuilder
	WithSignatures(signatures [][]byte) CompletedTaskBuilder
	Now() (CompletedTask, error)
}

// CompletedTask represents a completed task
type CompletedTask interface {
	Hash() hash.Hash
	Task() Task
	Signatures() [][]byte
}
