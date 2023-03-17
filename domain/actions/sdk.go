package actions

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias"
	"github.com/steve-care-software/propositions/domain/actions/resources"
	"github.com/steve-care-software/propositions/domain/users"
)

// NewActionBuilder creates a new action builder instance
func NewActionBuilder() ActionBuilder {
	hashAdapter := hash.NewAdapter()
	return createActionBuilder(hashAdapter)
}

// Builder represents an action builder
type Builder interface {
	Create() Builder
	WithList(list []Action) Builder
	Now() (Actions, error)
}

// Actions represents actions
type Actions interface {
	Hash() hash.Hash
	List() []Action
}

// ActionBuilder represents an action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithResource(resource resources.Resource) ActionBuilder
	WithMedia(media medias.Media) ActionBuilder
	WithParent(parent Action) ActionBuilder
	PostedBy(postedBy users.User) ActionBuilder
	CreatedOn(createdOn time.Time) ActionBuilder
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
