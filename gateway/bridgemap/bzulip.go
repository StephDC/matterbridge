// +build !nozulip

package bridgemap

import (
	bzulip "github.com/StephDC/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
