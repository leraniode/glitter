package palettes

import (
	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

var leraPetal = tone.New(
	tone.Light(80),
	tone.Vibrancy(83),
	tone.Hue(355),
	tone.Moody("playful"),
	tone.Named("lera.petal"),
)

var leraBloom = tone.New(
	tone.Light(73),
	tone.Vibrancy(99),
	tone.Hue(360),
	tone.Moody("playful"),
	tone.Named("lera.bloom"),
)

// lera is leraniode pink core identity tone.
var lera = tone.New(
	tone.Light(66),
	tone.Vibrancy(94),
	tone.Hue(3),
	tone.Moody("playful"),
	tone.Named("lera"),
)

var leraFlush = tone.New(
	tone.Light(56),
	tone.Vibrancy(96),
	tone.Hue(8),
	tone.Moody("playful"),
	tone.Named("lera.flush"),
)

var leraRose = tone.New(
	tone.Light(45),
	tone.Vibrancy(100),
	tone.Hue(12),
	tone.Moody("playful"),
	tone.Named("lera.rose"),
)

// LeraScale is the full Lera identity palette — the pink soul of Leraniode.
// Five tones from the softest petal to the deepest rose.
// Hue anchor: ~355–10° (warm pink, pulling toward magenta).
var LeraScale = register(palette.New("lera").
	Description("The lera identity palette. Pink spectrum of Leraniode.").
	Author("Leraniode").
	Version("0.1.0").
	Mood("playful").
	Add(leraPetal).
	Add(leraBloom).
	Add(lera).
	Add(leraFlush).
	Add(leraRose).
	MustBuild())
