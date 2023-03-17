package tasks

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
)

type builder struct {
	hashAdapter hash.Adapter
	quote       quotes.Quote
	managers    managers.Managers
	deposit     units.Unit
	pCreatedOn  *time.Time
	pExpiresOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		quote:       nil,
		managers:    nil,
		deposit:     nil,
		pCreatedOn:  nil,
		pExpiresOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithQuote adds a quote to the builder
func (app *builder) WithQuote(quote quotes.Quote) Builder {
	app.quote = quote
	return app
}

// WithManagers add managers to the builder
func (app *builder) WithManagers(managers managers.Managers) Builder {
	app.managers = managers
	return app
}

// WithDeposit adds deposit to the builder
func (app *builder) WithDeposit(deposit units.Unit) Builder {
	app.deposit = deposit
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// ExpiresOn adds an expiration time to the builder
func (app *builder) ExpiresOn(expiresOn time.Time) Builder {
	app.pExpiresOn = &expiresOn
	return app
}

// Now builds a new Task instance
func (app *builder) Now() (Task, error) {
	if app.quote == nil {
		return nil, errors.New("the quote is mandatory in order to build a Task instance")
	}

	if app.managers == nil {
		return nil, errors.New("the managers is mandatory in order to build a Task instance")
	}

	if app.deposit == nil {
		return nil, errors.New("the deposit is mandatory in order to build a Task instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Task instance")
	}

	if app.pExpiresOn == nil {
		return nil, errors.New("the expiration time is mandatory in order to build a Task instance")
	}

	if !app.pCreatedOn.Before(*app.pExpiresOn) {
		str := fmt.Sprintf("the creation time (%s) was expected to be before the expiration time (%s)", app.pCreatedOn.String(), app.pExpiresOn.String())
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.quote.Hash().Bytes(),
		app.managers.Hash().Bytes(),
		app.deposit.Hash().Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
		[]byte(fmt.Sprintf("%d", app.pExpiresOn.Unix())),
	})

	if err != nil {
		return nil, err
	}

	return createTask(
		*pHash,
		app.quote,
		app.managers,
		app.deposit,
		*app.pCreatedOn,
		*app.pExpiresOn,
	), nil
}
