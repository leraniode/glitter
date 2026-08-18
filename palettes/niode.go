package palettes

import (
	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

var niodeLavender = tone.New(
	tone.Light(69),
	tone.Vibrancy(80),
	tone.Hue(278),
	tone.Moody("vivid"),
	tone.Named("niode.lavender"),
)

var niodeLit = tone.New(
	tone.Light(63),
	tone.Vibrancy(100),
	tone.Hue(288),
	tone.Moody("vivid"),
	tone.Named("niode.lit"),
)

// niode is leraniode purple core identity tone.
var niode = tone.New(
	tone.Light(59),
	tone.Vibrancy(100),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("niode"),
)

var niodeBreeze = tone.New(
	tone.Light(49),
	tone.Vibrancy(86),
	tone.Hue(295),
	tone.Moody("vivid"),
	tone.Named("niode.breeze"),
)

var niodeDelta = tone.New(
	tone.Light(37),
	tone.Vibrancy(97),
	tone.Hue(292),
	tone.Moody("vivid"),
	tone.Named("niode.delta"),
)

// NiodeScale is the full Niode identity palette — the purple soul of Leraniode.
// Five tones from the coolest lavender to the deepest violet.
// Hue anchor: ~278–300° (violet, pulling toward blue-purple).
var NiodeScale = register(palette.New("niode").
	Description("The niode identity palette. Purple spectrum of Leraniode.").
	Author("Leraniode").
	Version("0.1.0").
	Mood("vivid").
	Add(niodeLavender).
	Add(niodeLit).
	Add(niode).
	Add(niodeBreeze).
	Add(niodeDelta).
	MustBuild())
