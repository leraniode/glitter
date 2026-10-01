package palettes

import (
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/wondertone/palette"
)

// Leraniode is the official base palette of the Leraniode organization.
// Org identity: the Spectrum arc (cyan through pink) — violet and pink are
// the org's own reserved range, never assigned to a product.
// Product tones: Green, Cyan, Blue, Indigo, Red — see ProductTones.
// Neutrals: a seven-step dark-only scale sharing the arc's violet hue.
var Leraniode = register(palette.New("leraniode").
	Description("The official Leraniode base palette. Foundation for all Leraniode products.").
	Author("Leraniode").
	Version("0.3.0").

	// — Org identity arc —
	Add(spectrumCyan).
	Add(spectrumBlue).
	Add(spectrumIndigo).
	Add(spectrumViolet).
	Add(spectrumPink).

	// — Product tones (extend the same arc into unreserved territory) —
	Add(tones.Red).
	Add(tones.Amber).
	Add(tones.Olive).
	Add(tones.Moss).
	Add(tones.Green).

	// — Neutral scale —
	Add(tones.Page).
	Add(tones.Panel).
	Add(tones.Surface).
	Add(tones.Border).
	Add(tones.Muted).
	Add(tones.Secondary).
	Add(tones.Primary).
	MustBuild())
