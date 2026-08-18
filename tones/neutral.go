package tones

import "github.com/leraniode/wondertone/tone"

// Neutral scale — violet-tinted, dark-first.
// Darker than Catppuccin Mocha. Designed for Leraniode's deep visual identity.
// Hue anchor: ~284° (violet-leaning neutral).

// Ink is the primary text tone — near-white with violet warmth.
var Ink = register(tone.New(
	tone.Light(82),
	tone.Vibrancy(45),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("ink"),
))

// Mute is secondary text — slightly dimmed, still readable.
var Mute = register(tone.New(
	tone.Light(70),
	tone.Vibrancy(30),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("mute"),
))

// Ash is tertiary text — visibly quieter, for metadata and captions.
var Ash = register(tone.New(
	tone.Light(61),
	tone.Vibrancy(22),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("ash"),
))

// Dim is the top overlay — for subdued interactive states.
var Dim = register(tone.New(
	tone.Light(53),
	tone.Vibrancy(17),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("dim"),
))

// Haze is mid overlay — separators, placeholder text.
var Haze = register(tone.New(
	tone.Light(45),
	tone.Vibrancy(12),
	tone.Hue(297),
	tone.Moody("focused"),
	tone.Named("haze"),
))

// Ghost is low overlay — faintest visible presence.
var Ghost = register(tone.New(
	tone.Light(35),
	tone.Vibrancy(15),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("ghost"),
))

// Rim is the highest surface — raised elements, hover states.
var Rim = register(tone.New(
	tone.Light(27),
	tone.Vibrancy(17),
	tone.Hue(297),
	tone.Moody("vivid"),
	tone.Named("rim"),
))

// Shell is mid surface — cards, input backgrounds.
var Shell = register(tone.New(
	tone.Light(21),
	tone.Vibrancy(19),
	tone.Hue(297),
	tone.Moody("urgent"),
	tone.Named("shell"),
))

// Veil is the lowest surface — subtle panel backgrounds.
var Veil = register(tone.New(
	tone.Light(17),
	tone.Vibrancy(21),
	tone.Hue(297),
	tone.Moody("urgent"),
	tone.Named("veil"),
))

// Slate is the page base — the main background color.
var Slate = register(tone.New(
	tone.Light(15),
	tone.Vibrancy(23),
	tone.Hue(297),
	tone.Moody("urgent"),
	tone.Named("slate"),
))

// Deep is below the base — panels, sidebars, drawers.
var Deep = register(tone.New(
	tone.Light(13),
	tone.Vibrancy(27),
	tone.Hue(297),
	tone.Moody("urgent"),
	tone.Named("deep"),
))

// Abyss is the darkest neutral — outermost chrome, true depth.
var Abyss = register(tone.New(
	tone.Light(11),
	tone.Vibrancy(30),
	tone.Hue(297),
	tone.Moody("urgent"),
	tone.Named("abyss"),
))
