// +build !notelegram

package bridgemap

import (
	btelegram "github.com/StephDC/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
