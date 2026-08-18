// Package tones contains standalone tones ready for use across Leraniode products.
//
// Tones are registered at package init via register() in the sibling files
// (neutral.go, warm.go, cool.go). Query and transform them with the utils package.
//
//	import (
//	    "github.com/leraniode/glitter/tones"
//	    "github.com/leraniode/glitter/utils"
//	)
//
//	ember, _ := utils.ToneByName(tones.All(), "ember")
//	dark := utils.Dark(tones.All())
package tones

import "github.com/leraniode/wondertone/tone"

var all []tone.Tone

func register(t tone.Tone) tone.Tone {
	all = append(all, t)
	return t
}

// All returns every registered tone in registration order.
func All() []tone.Tone {
	return all
}
