package quotes

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/users"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a quote builder
type Builder interface {
	Create() Builder
	WithProposition(proposition propositions.Proposition) Builder
	WithTitle(title string) Builder
	WithDescription(description string) Builder
	WithAmount(amount uint) Builder
	WithDuration(duration time.Duration) Builder
	CreatedOn(createdOn time.Time) Builder
	SubmittedBy(submittedBy users.User) Builder
	Now() (Quote, error)
}

// Quote represents a quote on a proposition
type Quote interface {
	Hash() hash.Hash
	Proposition() propositions.Proposition
	SubmittedBy() users.User
	Title() string
	Description() string
	Amount() uint
	CreatedOn() time.Time
	Duration() time.Duration
}
