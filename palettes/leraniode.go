package palettes

import (
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/wondertone/palette"
)

// Leraniode is the official base palette of the Leraniode organization.
// Identity: the Spectrum arc (cyan through pink, one continuous line).
// Accent: Confirmed — the one deliberate exception, reserved for proven states.
// Neutrals: a seven-step dark-only scale sharing the arc's violet hue.
var Leraniode = register(palette.New("leraniode").
	Description("The official Leraniode base palette. Foundation for all Leraniode products.").
	Author("Leraniode").
	Version("0.2.0").

	// — Identity arc —
	Add(spectrumCyan).
	Add(spectrumBlue).
	Add(spectrumIndigo).
	Add(spectrumViolet).
	Add(spectrumPink).

	// — Deliberate exception —
	Add(tones.Confirmed).

	// — Neutral scale —
	Add(tones.Page).
	Add(tones.Panel).
	Add(tones.Surface).
	Add(tones.Border).
	Add(tones.Muted).
	Add(tones.Secondary).
	Add(tones.Primary).
	MustBuild())
