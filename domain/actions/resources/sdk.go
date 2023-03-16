package resources

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	propositions "github.com/steve-care-software/propositions/domain"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
	"github.com/steve-care-software/propositions/domain/quotes"
	"github.com/steve-care-software/propositions/domain/quotes/managers"
	"github.com/steve-care-software/propositions/domain/quotes/managers/tasks"
)

// Builder represents a resource builder
type Builder interface {
	Create() Builder
	WithProposition(proposition propositions.Proposition) Builder
	WithQuote(quote quotes.Quote) Builder
	WithManager(manager managers.Manager) Builder
	WithTask(task tasks.Task) Builder
	WithComment(comment comments.Comment) Builder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	IsProposition() bool
	Proposition() propositions.Proposition
	IsQuote() bool
	Quote() quotes.Quote
	IsManager() bool
	Manager() managers.Manager
	IsTask() bool
	Task() tasks.Task
	IsComment() bool
	Comment() comments.Comment
}
