package resources

import (
	"errors"

	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

type builder struct {
	proposition propositions.Proposition
	quote       quotes.Quote
	manager     managers.Manager
	task        tasks.Task
	comment     comments.Comment
}

func createBuilder() Builder {
	out := builder{
		proposition: nil,
		quote:       nil,
		manager:     nil,
		task:        nil,
		comment:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithProposition adds a proposition to the builder
func (app *builder) WithProposition(proposition propositions.Proposition) Builder {
	app.proposition = proposition
	return app
}

// WithQuote adds a quote to the builder
func (app *builder) WithQuote(quote quotes.Quote) Builder {
	app.quote = quote
	return app
}

// WithManager adds a manager to the builder
func (app *builder) WithManager(manager managers.Manager) Builder {
	app.manager = manager
	return app
}

// WithTask adds a task to the builder
func (app *builder) WithTask(task tasks.Task) Builder {
	app.task = task
	return app
}

// WithComment adds a comment to the builder
func (app *builder) WithComment(comment comments.Comment) Builder {
	app.comment = comment
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.proposition != nil {

	}

	if app.quote != nil {

	}

	if app.manager != nil {

	}

	if app.task != nil {

	}

	if app.comment != nil {

	}

	return nil, errors.New("the Resource is invalid")
}
