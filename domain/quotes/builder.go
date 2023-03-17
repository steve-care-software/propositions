package quotes

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/users"
)

type builder struct {
	hashAdapter hash.Adapter
	proposition propositions.Proposition
	submittedBy users.User
	title       string
	description string
	amount      uint
	pCreatedOn  *time.Time
	duration    time.Duration
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		proposition: nil,
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
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithProposition adds a proposition to the builder
func (app *builder) WithProposition(proposition propositions.Proposition) Builder {
	app.proposition = proposition
	return app
}

// WithTitle adds a title to the builder
func (app *builder) WithTitle(title string) Builder {
	app.title = title
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithAmount adds an amount to the builder
func (app *builder) WithAmount(amount uint) Builder {
	app.amount = amount
	return app
}

// WithDuration adds a duration to the builder
func (app *builder) WithDuration(duration time.Duration) Builder {
	app.duration = duration
	return app
}

// SubmittedBy adds a submittedBy user to the builder
func (app *builder) SubmittedBy(submittedBy users.User) Builder {
	app.submittedBy = submittedBy
	return app
}

// CreatedOn adds a creation to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Quote instance
func (app *builder) Now() (Quote, error) {
	if app.proposition == nil {
		return nil, errors.New("the proposition is mandatory in order to build a Quote instance")
	}

	if app.title == "" {
		return nil, errors.New("the title is mandatory in order to build a Quote instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build a Quote instance")
	}

	if app.duration <= 0 {
		return nil, errors.New("the duration must be greater than zero (0) in order to build a Quote instance")
	}

	if app.submittedBy == nil {
		return nil, errors.New("the submittedBy user is mandatory in order to build a Quote instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Quote instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.proposition.Hash().Bytes(),
		app.submittedBy.Hash().Bytes(),
		[]byte(app.title),
		[]byte(fmt.Sprintf("%d", app.amount)),
		[]byte(fmt.Sprintf("%d", app.duration)),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	})

	if err != nil {
		return nil, err
	}

	return createQuote(
		*pHash,
		app.proposition,
		app.submittedBy,
		app.title,
		app.description,
		app.amount,
		*app.pCreatedOn,
		app.duration,
	), nil
}
