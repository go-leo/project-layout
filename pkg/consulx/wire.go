package consulx

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewClients,
)
