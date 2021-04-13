package palette

type Palette interface {
	LUT() [256][3]float32
	Lookup(x uint8) [3]float32
}
