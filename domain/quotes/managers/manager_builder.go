package managers

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/users"
)

type managerBuilder struct {
	hashAdapter hash.Adapter
	quote       quotes.Quote
	submittedBy users.User
	title       string
	description string
	amount      uint
	pCreatedOn  *time.Time
	duration    time.Duration
}

func createManagerBuilder(
	hashAdapter hash.Adapter,
) ManagerBuilder {
	out := managerBuilder{
		hashAdapter: hashAdapter,
		quote:       nil,
		submittedBy: nil,
		title:       "",
		description: "",
		amount:      0,
		pCreatedOn:  nil,
		duration:    0,
	}

	return &out
}

// Create initializes the builder
func (app *managerBuilder) Create() ManagerBuilder {
	return createManagerBuilder(app.hashAdapter)
}

// WithQuote adds a quote to the builder
func (app *managerBuilder) WithQuote(quote quotes.Quote) ManagerBuilder {
	app.quote = quote
	return app
}

// WithTitle adds a title to the builder
func (app *managerBuilder) WithTitle(title string) ManagerBuilder {
	app.title = title
	return app
}

// WithDescription adds a description to the builder
func (app *managerBuilder) WithDescription(description string) ManagerBuilder {
	app.description = description
	return app
}

// WithAmount adds an amount to the builder
func (app *managerBuilder) WithAmount(amount uint) ManagerBuilder {
	app.amount = amount
	return app
}

// WithDuration adds a duration to the builder
func (app *managerBuilder) WithDuration(duration time.Duration) ManagerBuilder {
	app.duration = duration
	return app
}

// SubmittedBy adds a submittedBy user to the builder
func (app *managerBuilder) SubmittedBy(submittedBy users.User) ManagerBuilder {
	app.submittedBy = submittedBy
	return app
}

// CreatedOn adds a creation to the builder
func (app *managerBuilder) CreatedOn(createdOn time.Time) ManagerBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Manager instance
func (app *managerBuilder) Now() (Manager, error) {
	if app.quote == nil {
		return nil, errors.New("the quote is mandatory in order to build a Manager instance")
	}

	if app.title == "" {
		return nil, errors.New("the title is mandatory in order to build a Manager instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Manager instance")
	}

	if app.duration <= 0 {
		return nil, errors.New("the duration must be greater than zero (0) in order to build a Manager instance")
	}

	if app.submittedBy == nil {
		return nil, errors.New("the submittedBy user is mandatory in order to build a Manager instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Manager instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.quote.Hash().Bytes(),
		app.submittedBy.Hash().Bytes(),
		[]byte(app.title),
		[]byte(fmt.Sprintf("%d", app.amount)),
		[]byte(fmt.Sprintf("%d", app.duration)),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	})

	if err != nil {
		return nil, err
	}

	return createManager(
		*pHash,
		app.quote,
		app.submittedBy,
		app.title,
		app.description,
		app.amount,
		*app.pCreatedOn,
		app.duration,
	), nil
}
