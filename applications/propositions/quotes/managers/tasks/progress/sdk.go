package progress

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks/progress"
)

// Builder represents a progress application builder
type Builder interface {
	Create() Builder
	WithTask(task tasks.Task) Builder
	Now() (Application, error)
}

// Application represents a progress application
type Application interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (progress.Progress, error)
}
