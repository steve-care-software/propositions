package quotes

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/users"
)

type quote struct {
	hash        hash.Hash
	proposition propositions.Proposition
	submittedBy users.User
	title       string
	description string
	amount      uint
	createdOn   time.Time
	duration    time.Duration
}

func createQuote(
	hash hash.Hash,
	proposition propositions.Proposition,
	submittedBy users.User,
	title string,
	description string,
	amount uint,
	createdOn time.Time,
	duration time.Duration,
) Quote {
	out := quote{
		hash:        hash,
		proposition: proposition,
		submittedBy: submittedBy,
		title:       title,
		description: description,
		amount:      amount,
		createdOn:   createdOn,
		duration:    duration,
	}

	return &out
}

// Hash returns the hash
func (obj *quote) Hash() hash.Hash {
	return obj.hash
}

// Proposition returns the proposition
func (obj *quote) Proposition() propositions.Proposition {
	return obj.proposition
}

// SubmittedBy returns the submitted by user
func (obj *quote) SubmittedBy() users.User {
	return obj.submittedBy
}

// Title returns the title
func (obj *quote) Title() string {
	return obj.title
}

// Description returns the description
func (obj *quote) Description() string {
	return obj.description
}

// Amount returns the amount
func (obj *quote) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *quote) CreatedOn() time.Time {
	return obj.createdOn
}

// Duration returns the duration
func (obj *quote) Duration() time.Duration {
	return obj.duration
}
