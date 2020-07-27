// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/StephDC/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
