package actions

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias"
	"github.com/steve-care-software/propositions/domain/actions/resources"
	"github.com/steve-care-software/propositions/domain/users"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents an action builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource) Builder
	WithMedia(media medias.Media) Builder
	WithParent(parent Action) Builder
	PostedBy(postedBy users.User) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Action, error)
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	Resource() resources.Resource
	Media() medias.Media
	PostedBy() users.User
	CreatedOn() time.Time
	HasParent() bool
	Parent() Action
}
