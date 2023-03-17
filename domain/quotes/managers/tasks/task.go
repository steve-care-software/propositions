package tasks

import (
	"time"

	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
)

type task struct {
	hash      hash.Hash
	quote     quotes.Quote
	managers  managers.Managers
	deposit   units.Unit
	createdOn time.Time
	expiresOn time.Time
}

func createTask(
	hash hash.Hash,
	quote quotes.Quote,
	managers managers.Managers,
	deposit units.Unit,
	createdOn time.Time,
	expiresOn time.Time,
) Task {
	out := task{
		hash:      hash,
		quote:     quote,
		managers:  managers,
		deposit:   deposit,
		createdOn: createdOn,
		expiresOn: expiresOn,
	}

	return &out
}

// Hash returns the hash
func (obj *task) Hash() hash.Hash {
	return obj.hash
}

// Quote returns the quote
func (obj *task) Quote() quotes.Quote {
	return obj.quote
}

// Managers returns the managers
func (obj *task) Managers() managers.Managers {
	return obj.managers
}

// Deposit returns the deposit
func (obj *task) Deposit() units.Unit {
	return obj.deposit
}

// CreatedOn returns the creation time
func (obj *task) CreatedOn() time.Time {
	return obj.createdOn
}

// ExpiresOn returns the expiration time
func (obj *task) ExpiresOn() time.Time {
	return obj.expiresOn
}
