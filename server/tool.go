//go:build tools
// +build tools

package server

import (
	// wire dependencies
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/google/wire/internal/wire"
)
