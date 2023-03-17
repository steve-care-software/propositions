package tasks

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/applications/propositions/quotes/managers/tasks/progress"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

// Builder represents a task application builder
type Builder interface {
	Create() Builder
	WithManagers(managers managers.Managers) Builder
	Now() (Application, error)
}

// Application represents a task application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (tasks.Task, error)
	Progress(task tasks.Task) (progress.Application, error)
}
