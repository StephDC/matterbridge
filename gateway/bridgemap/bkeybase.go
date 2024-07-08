// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/StephDC/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
