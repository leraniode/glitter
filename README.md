<p align="center">
    <img width="1024" src="https://raw.githubusercontent.com/leraniode/.github/main/images/leraniode.svg" />
</p>

# Glitter

Official Leraniode tones and palettes gallery, built on [wondertone](https://github.com/leraniode/wondertone).

## Install

```bash
go get github.com/leraniode/glitter
```

## Packages

### `tones`

Standalone accent and neutral tone definitions.

```go
import "github.com/leraniode/glitter/tones"

for _, t := range tones.All() {
    fmt.Println(t.Name(), t.Hex())
}
```

### `palettes`

Available palettes:

| Palette      | Description                                    |
| ------------ | ---------------------------------------------- |
| `Leraniode`  | Full base palette (cores + accents + neutrals) |
| `LeraScale`  | Pink identity scale (5 tones)                  |
| `NiodeScale` | Purple identity scale (5 tones)                |
| `Clarity`    | Cyan → blue gradient scale (5 tones)           |

```go
import "github.com/leraniode/glitter/palettes"

lera := palettes.LeraScale.MustGet("lera")
for _, p := range palettes.All() {
    fmt.Println(p.Name(), p.Len())
}
```

### `utils`

Query, transform, and **export** tones/palettes.

#### Lookup & filter

```go
import (
    "github.com/leraniode/glitter/tones"
    "github.com/leraniode/glitter/utils"
)

ember, _ := utils.ToneByName(tones.All(), "ember")
dark := utils.Dark(tones.All())
warm := utils.ByMood(tones.All(), "warm")
nearRed := utils.ByHueRange(tones.All(), 350, 30)
```

Helpers: `ToneByName`, `MustToneByName`, `HasTone`, `First`, `Last`, `Skip`, `Range`, `At`,  
`PaletteByName`, `MustPaletteByName`, `HasPalette`, `PaletteNames`,  
`Filter`, `ByMood`, `ByLightRange`, `ByHueRange`, `Light`, `Dark`, `Names`, `Map`, `Find`.

#### Build & transform

```go
merged := utils.MergeTones(a, b)
gradient, _ := utils.GradientBetween(palettes.Leraniode.All(), "niode", "lera", 10)
p, _ := utils.BuildPalette("warm", utils.ByMood(tones.All(), "warm"),
    utils.WithDescription("warm accents only"),
)
```

Helpers: `MergeTones`, `RenamePrefix`, `WithEnergyAll`, `GradientBetween`, `MustGradientBetween`,  
`BuildPalette`, `MustBuildPalette`, `GoSource`.

#### Export: CSS / SCSS / Tailwind / JSON

```go
css := utils.CSSVariables(palettes.Leraniode.All(), utils.CSSOptions{
    Prefix: "--lera-",
    Format: "oklch", // or "hex", "rgb"
})

scss := utils.SCSSMap(palettes.LeraScale.All(), "lera", utils.CSSOptions{})
tw := utils.TailwindTheme(palettes.Leraniode.All(), utils.CSSOptions{})

jsonBytes, _ := utils.JSON(palettes.Clarity.All())
jsonPal, _ := utils.JSONPalette(palettes.Leraniode)
```

---

## License

MIT — Leraniode

---

Part of [Leraniode](https://github.com/leraniode).

<p align="center">
    <br/>
    <a href="https://github.com/leraniode">
        <img src="https://raw.githubusercontent.com/leraniode/.github/main/assets/footer.svg" width="1024" />
    </a>
</p>
