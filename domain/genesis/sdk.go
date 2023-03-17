package genesis

import (
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Treeshold() uint
}
