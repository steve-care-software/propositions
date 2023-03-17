package managers

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/users"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewManagerBuilder creates a new manager builder
func NewManagerBuilder() ManagerBuilder {
	hashAdapter := hash.NewAdapter()
	return createManagerBuilder(hashAdapter)
}

// Builder represents a manager's builder
type Builder interface {
	Create() Builder
	WithList(list []Manager) Builder
	Now() (Managers, error)
}

// Managers represents managers
type Managers interface {
	Hash() hash.Hash
	List() []Manager
}

// ManagerBuilder represents a manager ManagerBuilder
type ManagerBuilder interface {
	Create() ManagerBuilder
	WithQuote(quote quotes.Quote) ManagerBuilder
	WithTitle(title string) ManagerBuilder
	WithDescription(description string) ManagerBuilder
	WithAmount(amount uint) ManagerBuilder
	WithDuration(duration time.Duration) ManagerBuilder
	CreatedOn(createdOn time.Time) ManagerBuilder
	SubmittedBy(submittedBy users.User) ManagerBuilder
	Now() (Manager, error)
}

// Manager represents a manager's quote on a quote
type Manager interface {
	Hash() hash.Hash
	Quote() quotes.Quote
	SubmittedBy() users.User
	Title() string
	Description() string
	Amount() uint
	CreatedOn() time.Time
	Duration() time.Duration
}
