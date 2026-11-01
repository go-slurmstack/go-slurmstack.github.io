//go:build tools

// Package site keeps the landingkit requirement in go.mod: Hugo reads the
// module from there, and without an import `go mod tidy` (which Renovate
// runs) would delete it.
package site

import _ "github.com/go-fleettools/landingkit"
