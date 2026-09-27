package tones

import "github.com/leraniode/wondertone/tone"

// Confirmed is the one deliberate color kept outside the spectrum arc
// (see palettes/spectrum.go). It does not share the arc's hue/vibrancy
// recipe on purpose — it exists to be the exception, not another sample
// point. Reserved specifically for proven, verified, "this checked out"
// states. Not a general-purpose success/positive color — if a use doesn't
// mean "this was confirmed," it doesn't get this tone.
var Confirmed = register(tone.New(
	tone.Light(68),
	tone.Vibrancy(85),
	tone.Hue(150),
	tone.Moody("certain"),
	tone.Named("confirmed"),
))
