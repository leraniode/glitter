package palettes

import (
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/wondertone/mix"
	"github.com/leraniode/wondertone/tone"
)

// Deepen mixes t toward neutral.page in real OKLab space (the same math
// Gradient walks) and returns the result. Use this instead of hand-picking
// a darker Light/Vibrancy pair — a guessed dark tone tends to keep too much
// Vibrancy for how dark it is, since real colour desaturates as it nears
// black. amount is 0 (unchanged) to 1 (full neutral.page).
//
// This exists specifically so background treatments (wedges, washes) that
// need a "deep" version of an arc color are computed, not eyeballed.
func Deepen(t tone.Tone, amount float64) tone.Tone {
	return mix.Mix(t, tones.Page, amount)
}
