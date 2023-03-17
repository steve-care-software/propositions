package managers

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/users"
)

type manager struct {
	hash        hash.Hash
	quote       quotes.Quote
	submittedBy users.User
	title       string
	description string
	amount      uint
	createdOn   time.Time
	duration    time.Duration
}

func createManager(
	hash hash.Hash,
	quote quotes.Quote,
	submittedBy users.User,
	title string,
	description string,
	amount uint,
	createdOn time.Time,
	duration time.Duration,
) Manager {
	out := manager{
		hash:        hash,
		quote:       quote,
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
func (obj *manager) Hash() hash.Hash {
	return obj.hash
}

// Quote returns the quote
func (obj *manager) Quote() quotes.Quote {
	return obj.quote
}

// SubmittedBy returns the submitted by user
func (obj *manager) SubmittedBy() users.User {
	return obj.submittedBy
}

// Title returns the title
func (obj *manager) Title() string {
	return obj.title
}

// Description returns the description
func (obj *manager) Description() string {
	return obj.description
}

// Amount returns the amount
func (obj *manager) Amount() uint {
	return obj.amount
}

// CreatedOn returns the creation time
func (obj *manager) CreatedOn() time.Time {
	return obj.createdOn
}

// Duration returns the duration
func (obj *manager) Duration() time.Duration {
	return obj.duration
}
