package tones

import "github.com/leraniode/wondertone/tone"

// Neutral scale — dark-only, one recipe: hue and vibrancy stay fixed at every
// step, only Light changes. Hue anchor: 296° (matches spectrum.violet in
// palettes/spectrum.go, so neutrals are visibly the same identity, just
// desaturated) — not an unrelated scale invented on its own.
//
// Seven roles, darkest to lightest: page, panel, surface, border, muted,
// secondary, primary.

// Page is the page background — the darkest surface.
var Page = register(tone.New(
	tone.Light(8),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.page"),
))

// Panel is a raised surface one step above the page.
var Panel = register(tone.New(
	tone.Light(12),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.panel"),
))

// Surface is a card or input background.
var Surface = register(tone.New(
	tone.Light(17),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.surface"),
))

// Border is a hairline or dividing edge.
var Border = register(tone.New(
	tone.Light(52),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.border"),
))

// Muted is tertiary text — hints, metadata, timestamps.
var Muted = register(tone.New(
	tone.Light(55),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.muted"),
))

// Secondary is body-adjacent text — captions, descriptions.
var Secondary = register(tone.New(
	tone.Light(70),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.secondary"),
))

// Primary is the primary text color — near-white with a whisper of violet.
var Primary = register(tone.New(
	tone.Light(88),
	tone.Vibrancy(10),
	tone.Hue(296),
	tone.Moody("quiet"),
	tone.Named("neutral.primary"),
))
