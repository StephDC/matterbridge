// +build !nowhatsapp

package bridgemap

import (
	bwhatsapp "github.com/StephDC/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
