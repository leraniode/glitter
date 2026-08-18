// Package palettes provides the official Leraniode colour palettes built on
// wondertone.
package palettes

import (
	"github.com/leraniode/wondertone/palette"
)

var registry []*palette.Palette

func register(p *palette.Palette) *palette.Palette {
	registry = append(registry, p)
	return p
}

// All returns every registered palette in registration order.
func All() []*palette.Palette {
	return registry
}
