package users

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type user struct {
	hash      hash.Hash
	name      string
	pubKey    hash.Hash
	createdOn time.Time
}

func createUser(
	hash hash.Hash,
	name string,
	pubKey hash.Hash,
	createdOn time.Time,
) User {
	out := user{
		hash:      hash,
		name:      name,
		pubKey:    pubKey,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *user) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *user) Name() string {
	return obj.name
}

// PubKey returns the pubKey hash
func (obj *user) PubKey() hash.Hash {
	return obj.pubKey
}

// CreatedOn returns the creation time
func (obj *user) CreatedOn() time.Time {
	return obj.createdOn
}
