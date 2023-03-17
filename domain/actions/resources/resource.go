package resources

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

type resource struct {
	proposition propositions.Proposition
	quote       quotes.Quote
	manager     managers.Manager
	task        tasks.Task
	comment     comments.Comment
}

func createResourceWithProposition(
	proposition propositions.Proposition,
) Resource {
	return createResourceInternally(proposition, nil, nil, nil, nil)
}

func createResourceWithQuote(
	quote quotes.Quote,
) Resource {
	return createResourceInternally(nil, quote, nil, nil, nil)
}

func createResourceWithManager(
	manager managers.Manager,
) Resource {
	return createResourceInternally(nil, nil, manager, nil, nil)
}

func createResourceWithTask(
	task tasks.Task,
) Resource {
	return createResourceInternally(nil, nil, nil, task, nil)
}

func createResourceWithComment(
	comment comments.Comment,
) Resource {
	return createResourceInternally(nil, nil, nil, nil, comment)
}

func createResourceInternally(
	proposition propositions.Proposition,
	quote quotes.Quote,
	manager managers.Manager,
	task tasks.Task,
	comment comments.Comment,
) Resource {
	out := resource{
		proposition: proposition,
		quote:       quote,
		manager:     manager,
		task:        task,
		comment:     comment,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	if obj.IsProposition() {
		return obj.proposition.Hash()
	}

	if obj.IsQuote() {
		return obj.quote.Hash()
	}

	if obj.IsManager() {
		return obj.manager.Hash()
	}

	if obj.IsTask() {
		return obj.task.Hash()
	}

	return obj.comment.Hash()
}

// IsProposition returns true if there is a proposition, false otherwise
func (obj *resource) IsProposition() bool {
	return obj.proposition != nil
}

// Proposition returns the proposition, if any
func (obj *resource) Proposition() propositions.Proposition {
	return obj.proposition
}

// IsQuote returns true if there is a quote, false otherwise
func (obj *resource) IsQuote() bool {
	return obj.quote != nil
}

// Quote returns the quote, if any
func (obj *resource) Quote() quotes.Quote {
	return obj.quote
}

// IsManager returns true if there is a manager, false otherwise
func (obj *resource) IsManager() bool {
	return obj.manager != nil
}

// Manager returns the manager, if any
func (obj *resource) Manager() managers.Manager {
	return obj.manager
}

// IsTask returns true if there is a task, false otherwise
func (obj *resource) IsTask() bool {
	return obj.task != nil
}

// Task returns the task, if any
func (obj *resource) Task() tasks.Task {
	return obj.task
}

// IsComment returns true if there is a comment, false otherwise
func (obj *resource) IsComment() bool {
	return obj.comment != nil
}

// Comment returns the comment, if any
func (obj *resource) Comment() comments.Comment {
	return obj.comment
}
