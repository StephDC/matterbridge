// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/StephDC/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
