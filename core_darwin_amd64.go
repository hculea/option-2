//go:build darwin && amd64
// +build darwin,amd64

package test

import _ "embed"

//go:embed option-2
var Hello string
