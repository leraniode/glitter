package palettes

import (
	"github.com/leraniode/wondertone/palette"
	"github.com/leraniode/wondertone/tone"
)

var clarityTide = tone.New(
	tone.Light(90),
	tone.Vibrancy(100),
	tone.Hue(182),
	tone.Moody("warm"),
	tone.Named("clarity.tide"),
)

// Sky represents clarity in the leraniode palette
var claritySky = tone.New(
	tone.Light(80),
	tone.Vibrancy(100),
	tone.Hue(205),
	tone.Moody("open"),
	tone.Named("clarity.sky"),
)

var clarityMist = tone.New(
	tone.Light(70),
	tone.Vibrancy(100),
	tone.Hue(220),
	tone.Moody("mystical"),
	tone.Named("clarity.mist"),
)

var claritySapphire = tone.New(
	tone.Light(60),
	tone.Vibrancy(100),
	tone.Hue(238),
	tone.Moody("cool"),
	tone.Named("clarity.sapphire"),
)

var clarityDeep = tone.New(
	tone.Light(50),
	tone.Vibrancy(100),
	tone.Hue(258),
	tone.Moody("rich"),
	tone.Named("clarity.deep"),
)

// Clarity is the full Clarity gradient palette — the mystical glow of Leraniode.
// Five tones from warm teal-cyan through deep electric blue.
// Clarity is not one colour. It is a gradient lived across five tones.
// Hue anchor: ~182–262° (teal → sky → sapphire → blue → deep).
var Clarity = register(palette.New("clarity").
	Description("The clarity palette. Cyan-to-blue gradient of Leraniode.").
	Author("Leraniode").
	Version("0.1.0").
	Add(clarityTide).
	Add(claritySky).
	Add(clarityMist).
	Add(claritySapphire).
	Add(clarityDeep).
	MustBuild())
