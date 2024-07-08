// +build !novk

package bridgemap

import (
	bvk "github.com/StephDC/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
