package tasks

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

// Builder represents progress
type Builder interface {
	Create() Builder
	WithTask(task tasks.Task) Builder
	WithActions(actions actions.Actions) Builder
	Now() (Progress, error)
}

// Progress represents progress
type Progress interface {
	Hash() hash.Hash
	Task() tasks.Task
	Actions() actions.Actions
}
