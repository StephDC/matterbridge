//go:build !noharmony
// +build !noharmony

package bridgemap

import (
	bharmony "github.com/StephDC/matterbridge/bridge/harmony"
)

func init() {
	FullMap["harmony"] = bharmony.New
}
