package palettes

import (
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/wondertone/palette"
)

// Leraniode is the official base palette of the Leraniode organization.
// Identity: Lera (core pink) + Niode (core purple) + Clarity (blended cyan-blue).
// Accents: warm and cool tones with unique Leraniode names.
// Neutrals: violet-tinted, darker than standard palettes.
var Leraniode = register(palette.New("leraniode").
	Description("The official Leraniode base palette. Foundation for all Leraniode products.").
	Author("Leraniode").
	Version("0.1.0").

	// — Identity cores —
	Add(lera).
	Add(niode).
	Add(claritySky).

	// — Warm accents —
	Add(tones.Ember).
	Add(tones.Blaze).
	Add(tones.Dusk).
	Add(tones.Sol).
	Add(tones.Grove).

	// — Cool accents —
	Add(tones.Frost).
	Add(tones.Streak).
	Add(tones.Arc).
	Add(tones.Void).

	// — Neutral scale —
	Add(tones.Ink).
	Add(tones.Mute).
	Add(tones.Ash).
	Add(tones.Dim).
	Add(tones.Haze).
	Add(tones.Ghost).
	Add(tones.Rim).
	Add(tones.Shell).
	Add(tones.Veil).
	Add(tones.Slate).
	Add(tones.Deep).
	Add(tones.Abyss).
	MustBuild())
