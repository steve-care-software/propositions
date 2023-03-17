package actions

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias"
	"github.com/steve-care-software/propositions/domain/actions/resources"
	"github.com/steve-care-software/propositions/domain/users"
)

type action struct {
	hash      hash.Hash
	resource  resources.Resource
	media     medias.Media
	postedBy  users.User
	createdOn time.Time
	parent    Action
}

func createAction(
	hash hash.Hash,
	resource resources.Resource,
	media medias.Media,
	postedBy users.User,
	createdOn time.Time,
) Action {
	return createActionInternally(hash, resource, media, postedBy, createdOn, nil)
}

func createActionWithParent(
	hash hash.Hash,
	resource resources.Resource,
	media medias.Media,
	postedBy users.User,
	createdOn time.Time,
	parent Action,
) Action {
	return createActionInternally(hash, resource, media, postedBy, createdOn, parent)
}

func createActionInternally(
	hash hash.Hash,
	resource resources.Resource,
	media medias.Media,
	postedBy users.User,
	createdOn time.Time,
	parent Action,
) Action {
	out := action{
		hash:      hash,
		resource:  resource,
		media:     media,
		postedBy:  postedBy,
		createdOn: createdOn,
		parent:    parent,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// Resource returns the resource
func (obj *action) Resource() resources.Resource {
	return obj.resource
}

// Media returns the media
func (obj *action) Media() medias.Media {
	return obj.media
}

// PostedBy returns the postedBy user
func (obj *action) PostedBy() users.User {
	return obj.postedBy
}

// CreatedOn returns the creation time
func (obj *action) CreatedOn() time.Time {
	return obj.createdOn
}

// HasParent returns true if there is a aprent, false otherwise
func (obj *action) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *action) Parent() Action {
	return obj.parent
}
