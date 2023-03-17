package managers

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Manager
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Manager) Builder {
	app.list = list
	return app
}

// Now builds a new Managers instance
func (app *builder) Now() (Managers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Manager in order to build a Managers instance")
	}

	data := [][]byte{}
	for _, oneTrx := range app.list {
		data = append(data, oneTrx.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createManagers(*pHash, app.list), nil
}
