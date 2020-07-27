// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/StephDC/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
