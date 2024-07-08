// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/StephDC/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
