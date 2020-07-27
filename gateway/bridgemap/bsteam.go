// +build !nosteam

package bridgemap

import (
	bsteam "github.com/StephDC/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
