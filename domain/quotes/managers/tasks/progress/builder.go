package progress

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

type builder struct {
	hashAdapter hash.Adapter
	task        tasks.Task
	actions     actions.Actions
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		task:        nil,
		actions:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithTask adds a task to the builder
func (app *builder) WithTask(task tasks.Task) Builder {
	app.task = task
	return app
}

// WithActions add actions to the builder
func (app *builder) WithActions(actions actions.Actions) Builder {
	app.actions = actions
	return app
}

// Now builds a new Progresss instance
func (app *builder) Now() (Progress, error) {
	if app.task == nil {
		return nil, errors.New("the task is mandatory in order to build a Progress instance")
	}

	if app.actions == nil {
		return nil, errors.New("the actions is mandatory in order to build a Progress instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.task.Hash().Bytes(),
		app.actions.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createProgress(*pHash, app.task, app.actions), nil
}
