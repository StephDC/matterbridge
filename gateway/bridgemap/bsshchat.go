// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/StephDC/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
