// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/StephDC/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
