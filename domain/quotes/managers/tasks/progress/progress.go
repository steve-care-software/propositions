package progress

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

type progress struct {
	hash    hash.Hash
	task    tasks.Task
	actions actions.Actions
}

func createProgress(
	hash hash.Hash,
	task tasks.Task,
	actions actions.Actions,
) Progress {
	out := progress{
		hash:    hash,
		task:    task,
		actions: actions,
	}

	return &out
}

// Hash returns the hash
func (obj *progress) Hash() hash.Hash {
	return obj.hash
}

// Task returns the task
func (obj *progress) Task() tasks.Task {
	return obj.task
}

// Actions returns the actions
func (obj *progress) Actions() actions.Actions {
	return obj.actions
}
