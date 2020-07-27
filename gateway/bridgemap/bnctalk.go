// +build !nonctalk

package bridgemap

import (
	btalk "github.com/StephDC/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
