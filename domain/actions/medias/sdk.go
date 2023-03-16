package medias

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
	"github.com/steve-care-software/propositions/domain/actions/medias/votes"
)

// Builder represents a media builder
type Builder interface {
	Create() Builder
	WithVote(vote votes.Vote) Builder
	WithComment(comment comments.Comment) Builder
	Now() (Media, error)
}

// Media represents a media
type Media interface {
	Hash() hash.Hash
	IsVote() bool
	Vote() votes.Vote
	IsComment() bool
	Comment() comments.Comment
}
