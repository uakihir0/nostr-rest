//go:build tools
// +build tools

package server

import (
	// openapi codegen
	_ "gopkg.in/yaml.v2"
	// wire dependencies
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/google/wire/internal/wire"
)
