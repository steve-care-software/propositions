package actions

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias"
	"github.com/steve-care-software/propositions/domain/actions/resources"
	"github.com/steve-care-software/propositions/domain/users"
)

type actionActionBuilder struct {
	hashAdapter hash.Adapter
	resource    resources.Resource
	media       medias.Media
	postedBy    users.User
	pCreatedOn  *time.Time
	parent      Action
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionActionBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		media:       nil,
		postedBy:    nil,
		pCreatedOn:  nil,
		parent:      nil,
	}

	return &out
}

// Create initializes the actionActionBuilder
func (app *actionActionBuilder) Create() ActionBuilder {
	return createActionBuilder(app.hashAdapter)
}

// WithResource adds a resource to the actionActionBuilder
func (app *actionActionBuilder) WithResource(resource resources.Resource) ActionBuilder {
	app.resource = resource
	return app
}

// WithMedia adds a media to the actionActionBuilder
func (app *actionActionBuilder) WithMedia(media medias.Media) ActionBuilder {
	app.media = media
	return app
}

// WithParent adds a parent action to the actionActionBuilder
func (app *actionActionBuilder) WithParent(parent Action) ActionBuilder {
	app.parent = parent
	return app
}

// PostedBy adds a postedBy user to the actionActionBuilder
func (app *actionActionBuilder) PostedBy(postedBy users.User) ActionBuilder {
	app.postedBy = postedBy
	return app
}

// CreatedOn adds a creation time to the actionActionBuilder
func (app *actionActionBuilder) CreatedOn(createdOn time.Time) ActionBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Action instance
func (app *actionActionBuilder) Now() (Action, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Action instance")
	}

	if app.media == nil {
		return nil, errors.New("the media is mandatory in order to build an Action instance")
	}

	if app.postedBy == nil {
		return nil, errors.New("the postedBy user is mandatory in order to build an Action instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Action instance")
	}

	data := [][]byte{
		app.resource.Hash().Bytes(),
		app.media.Hash().Bytes(),
		app.postedBy.Hash().Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	}

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createActionWithParent(*pHash, app.resource, app.media, app.postedBy, *app.pCreatedOn, app.parent), nil
	}

	return createAction(*pHash, app.resource, app.media, app.postedBy, *app.pCreatedOn), nil
}
