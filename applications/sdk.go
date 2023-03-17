package applications

import (
	application_identities "github.com/steve-care-software/identities/applications"
	identities "github.com/steve-care-software/identities/domain"
	"github.com/steve-care-software/propositions/applications/propositions"
	"github.com/steve-care-software/propositions/applications/users"
)

// Application represents a proposition application
type Application interface {
	Identity() application_identities.Application
	Proposition(identity identities.Identity) propositions.Application
	User() users.Application
}
