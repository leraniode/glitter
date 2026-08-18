package tones

import "github.com/leraniode/wondertone/tone"

// Cool accent tones.

// Frost is a sharp, icy cyan-white — cold clarity at its edge.
var Frost = register(tone.New(
	tone.Light(80),
	tone.Vibrancy(64),
	tone.Hue(195),
	tone.Moody("cool"),
	tone.Named("frost"),
))

var Streak = register(tone.New(
	tone.Light(70),
	tone.Vibrancy(68),
	tone.Hue(230),
	tone.Moody("mystical"),
	tone.Named("streak"),
))

// Arc is a vivid electric blue — charged, reactive.
var Arc = register(tone.New(
	tone.Light(60),
	tone.Vibrancy(100),
	tone.Hue(250),
	tone.Moody("electric"),
	tone.Named("arc"),
))

// Void is a deep, near-indigo blue — authority and depth.
var Void = register(tone.New(
	tone.Light(50),
	tone.Vibrancy(62),
	tone.Hue(270),
	tone.Moody("deep"),
	tone.Named("void"),
))
