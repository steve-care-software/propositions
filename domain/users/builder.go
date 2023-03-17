package users

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	pubKey      hash.Hash
	pCreatedOn  *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		pubKey:      nil,
		pCreatedOn:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPubKey adds a pubKey to the builder
func (app *builder) WithPubKey(pubKey hash.Hash) Builder {
	app.pubKey = pubKey
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new User instance
func (app *builder) Now() (User, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a User instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the pubKey is mandatory in order to build a User instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a User instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.pubKey.Bytes(),
		[]byte(fmt.Sprintf("%d", app.pCreatedOn.Unix())),
	})

	if err != nil {
		return nil, err
	}

	return createUser(*pHash, app.name, app.pubKey, *app.pCreatedOn), nil
}
