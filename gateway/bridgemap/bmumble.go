// +build !nomumble

package bridgemap

import (
	bmumble "github.com/StephDC/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
