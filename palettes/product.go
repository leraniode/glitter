package palettes

import (
	"github.com/leraniode/glitter/tones"
	"github.com/leraniode/wondertone/tone"
)

// ProductTones is the ordered set of tones a product may be assigned —
// never picked by eye, never a free color choice. Violet and pink are
// never in this list: they're the org mark's own reserved range.
// Everything else on the arc's circle, going red -> amber -> olive ->
// moss -> green -> cyan -> blue -> indigo, is product territory.
//
// Reusing a tone's underlying hue in more than one context is fine — the
// arc's cyan is the same cyan whether it's read as "part of the org
// Spectrum" or "assigned to Domtea." What matters is how a tone is used
// in a given place, not whether the same value appears elsewhere too —
// two products may share a tone when their names, context, and usage
// keep them legible as different things.
var ProductTones = []tone.Tone{
	tones.Red,      // 20°
	tones.Amber,    // 53°
	tones.Olive,    // 86°
	tones.Moss,     // 119°
	tones.Green,    // 152°
	spectrumCyan,   // 190°
	spectrumBlue,   // 228°
	spectrumIndigo, // 266°
}

// ProductAccent returns the product tone at position i. There are
// currently 8 slots. Running out again means the available 322° range
// itself needs a real decision — not another silent interpolation.
func ProductAccent(i int) tone.Tone {
	return ProductTones[i]
}
