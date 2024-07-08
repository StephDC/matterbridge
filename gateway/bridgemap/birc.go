// +build !noirc

package bridgemap

import (
	birc "github.com/StephDC/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
