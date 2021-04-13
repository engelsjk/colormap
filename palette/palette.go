package palette

type Palette interface {
	LUT() [256][3]float32
	Lookup(x uint8) [3]float32
}

// compile time checks
var (
	_ Palette = Cividis{}
	_ Palette = Crest{}
	_ Palette = Flare{}
	_ Palette = Icefire{}
	_ Palette = Inferno{}
	_ Palette = Magma{}
	_ Palette = Mako{}
	_ Palette = Plasma{}
	_ Palette = Rocket{}
	_ Palette = Turbo{}
	_ Palette = Viridis{}
	_ Palette = Vlag{}
)
