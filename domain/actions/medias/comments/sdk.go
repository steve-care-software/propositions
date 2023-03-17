package comments

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a comment builder
type Builder interface {
	Create() Builder
	WithContent(content string) Builder
	Now() (Comment, error)
}

// Comment represents a comment
type Comment interface {
	Hash() hash.Hash
	Content() string
}
