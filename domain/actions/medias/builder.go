package medias

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
	"github.com/steve-care-software/propositions/domain/actions/medias/comments"
)

type builder struct {
	hashAdapter hash.Adapter
	pVote       *uint8
	comment     comments.Comment
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pVote:       nil,
		comment:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithVote adds a vote to the builder
func (app *builder) WithVote(vote uint8) Builder {
	app.pVote = &vote
	return app
}

// WithComment adds a comment to the builder
func (app *builder) WithComment(comment comments.Comment) Builder {
	app.comment = comment
	return app
}

// Now builds a new Media instance
func (app *builder) Now() (Media, error) {
	data := [][]byte{}
	if app.pVote != nil {
		data = append(data, []byte(fmt.Sprintf("%d", *app.pVote)))
	}

	if app.comment != nil {
		data = append(data, app.comment.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Media is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pVote != nil {
		return createMediaWithVote(*pHash, app.pVote), nil
	}

	return createMediaWithComment(*pHash, app.comment), nil
}
