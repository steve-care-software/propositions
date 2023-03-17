package comments

import "github.com/steve-care-software/libs/cryptography/hash"

type comment struct {
	hash    hash.Hash
	content string
}

func createComment(
	hash hash.Hash,
	content string,
) Comment {
	out := comment{
		hash:    hash,
		content: content,
	}

	return &out
}

// Hash returns the hash
func (obj *comment) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *comment) Content() string {
	return obj.content
}
