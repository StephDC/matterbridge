// +build !noapi

package bridgemap

import (
	"github.com/StephDC/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
