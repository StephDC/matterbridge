// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/StephDC/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
