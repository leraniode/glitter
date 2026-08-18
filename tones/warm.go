package tones

import "github.com/leraniode/wondertone/tone"

// Warm accent tones — fire, earth, bloom energy.
// These are standalone accents, not identity tones.
// For lera (pink identity), see palettes/lera.go.

// Ember is a deep, saturated red — heat and urgency.
var Ember = register(tone.New(
	tone.Light(62),
	tone.Vibrancy(97),
	tone.Hue(25),
	tone.Moody("intense"),
	tone.Named("ember"),
))

// Blaze is a warm orange-red — energy between ember and dusk.
var Blaze = register(tone.New(
	tone.Light(60),
	tone.Vibrancy(100),
	tone.Hue(38),
	tone.Moody("warm"),
	tone.Named("blaze"),
))

// Dusk is a rich amber-orange — the color of last light.
var Dusk = register(tone.New(
	tone.Light(74),
	tone.Vibrancy(100),
	tone.Hue(55),
	tone.Moody("warm"),
	tone.Named("dusk"),
))

// Sol is a deep golden yellow — vivid, sun-core energy.
var Sol = register(tone.New(
	tone.Light(82),
	tone.Vibrancy(100),
	tone.Hue(88),
	tone.Moody("bright"),
	tone.Named("sol"),
))

// Grove is a rich, saturated green — alive, deep forest.
var Grove = register(tone.New(
	tone.Light(76),
	tone.Vibrancy(80),
	tone.Hue(142),
	tone.Moody("natural"),
	tone.Named("grove"),
))
