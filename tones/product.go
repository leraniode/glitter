package tones

import "github.com/leraniode/wondertone/tone"

// Green and Red extend the same +38° sequence the org's Spectrum arc
// already runs on (cyan 190, blue 228, indigo 266, violet 304, pink 342) —
// one step before cyan, and one step past pink, wrapping the circle.
// Nothing new was invented to produce them: the formula that already
// existed just wasn't stopped at the old ends.
//
// Amber, Olive, and Moss fill the one real gap left in the product
// range — the 132° span between Red (20°) and Green (152°), which the
// arc's own +38° step doesn't divide evenly (it overshoots to 172°, just
// 18° from Green, half the normal spacing everywhere else). Rather than
// force a step that doesn't fit, this gap is divided evenly on its own
// terms: 132° / 4 = 33° per step, giving three new interior tones with
// uniform spacing and no crowding.
//
// All extension tones share the arc's own recipe (Light 64, Vibrancy 92)
// so a product assigned any of them reads as "arc," not as a special case.
var Green = register(tone.New(
	tone.Light(64), tone.Vibrancy(92), tone.Hue(152), // 190 - 38
	tone.Moody("clear"), tone.Named("tone.green"),
))

var Red = register(tone.New(
	tone.Light(64), tone.Vibrancy(92), tone.Hue(20), // 342 + 38, wraps
	tone.Moody("clear"), tone.Named("tone.red"),
))

var Amber = register(tone.New(
	tone.Light(64), tone.Vibrancy(92), tone.Hue(53), // 20 + 33
	tone.Moody("clear"), tone.Named("tone.amber"),
))

var Olive = register(tone.New(
	tone.Light(64), tone.Vibrancy(92), tone.Hue(86), // 20 + 66
	tone.Moody("clear"), tone.Named("tone.olive"),
))

var Moss = register(tone.New(
	tone.Light(64), tone.Vibrancy(92), tone.Hue(119), // 20 + 99
	tone.Moody("clear"), tone.Named("tone.moss"),
))
