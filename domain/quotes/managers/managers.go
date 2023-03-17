package managers

import "github.com/steve-care-software/libs/cryptography/hash"

type managers struct {
	hash hash.Hash
	list []Manager
}

func createManagers(
	hash hash.Hash,
	list []Manager,
) Managers {
	out := managers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *managers) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *managers) List() []Manager {
	return obj.list
}
