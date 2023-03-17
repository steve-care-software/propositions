package medias

import (
	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
)

type media struct {
	hash    hash.Hash
	pVote   *uint8
	comment comments.Comment
}

func createMediaWithVote(
	hash hash.Hash,
	pVote *uint8,
) Media {
	return createMediaInternally(hash, pVote, nil)
}

func createMediaWithComment(
	hash hash.Hash,
	comment comments.Comment,
) Media {
	return createMediaInternally(hash, nil, comment)
}

func createMediaInternally(
	hash hash.Hash,
	pVote *uint8,
	comment comments.Comment,
) Media {
	out := media{
		hash:    hash,
		pVote:   pVote,
		comment: comment,
	}

	return &out
}

// Hash returns the hash
func (obj *media) Hash() hash.Hash {
	return obj.hash
}

// IsVote returns true if there is a vote, false otherwise
func (obj *media) IsVote() bool {
	return obj.pVote != nil
}

// Vote returns the vote, if any
func (obj *media) Vote() *uint8 {
	return obj.pVote
}

// IsComment returns true if there is a comment, false otherwise
func (obj *media) IsComment() bool {
	return obj.comment != nil
}

// Comment returns the comment, if any
func (obj *media) Comment() comments.Comment {
	return obj.comment
}
