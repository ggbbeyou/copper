package cauth

import (
	"github.com/google/wire"
)

// WireModule can be used as part of google/wire setup.
var WireModule = wire.NewSet( // nolint:gochecknoglobals
	NewRepo,
	NewSvc,
	NewVerifySessionMiddleware,

	wire.Struct(new(NewRouterParams), "*"),
	NewRouter,
)
