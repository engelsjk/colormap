package palette

type Empty struct{}

func (e Empty) LUT() [256][3]float32 {
	return [256][3]float32{}
}

func (e Empty) Lookup(x uint8) [3]float32 {
	return e.LUT()[x]
}
