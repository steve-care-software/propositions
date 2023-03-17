package comments

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content string) Builder {
	app.content = content
	return app
}

// Now builds a new Comment instance
func (app *builder) Now() (Comment, error) {
	if app.content == "" {
		return nil, errors.New("the content is mandatory in order to build a Comment instance")
	}

	pHash, err := app.hashAdapter.FromString(app.content)
	if err != nil {
		return nil, err
	}

	return createComment(*pHash, app.content), nil
}
