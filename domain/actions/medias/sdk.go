package medias

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
)

const (
	// VoteIsAccepted represents an accepted vote
	VoteIsAccepted (uint8) = iota

	// VoteIsDenied represents a denied vote
	VoteIsDenied
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a media builder
type Builder interface {
	Create() Builder
	WithVote(vote uint8) Builder
	WithComment(comment comments.Comment) Builder
	Now() (Media, error)
}

// Media represents a media
type Media interface {
	Hash() hash.Hash
	IsVote() bool
	Vote() *uint8
	IsComment() bool
	Comment() comments.Comment
}
