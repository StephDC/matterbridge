// +build !nogitter

package bridgemap

import (
	bgitter "github.com/StephDC/matterbridge/bridge/gitter"
)

func init() {
	FullMap["gitter"] = bgitter.New
}
