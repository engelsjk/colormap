package colormap

// 'colormap' is a ported copy of mikhailov-work/turbo_colormap.py
// https://gist.github.com/mikhailov-work/ee72ba4191942acecc03fe6da94fc73f
// Copyright 2019 Google LLC.
// SPDX-License-Identifier: Apache-2.0
// Author: Anton Mikhailov

import (
	"fmt"
	"image/color"
	"math"

	"github.com/engelsjk/colormap/palette"
)

func Palettes() []string {
	return []string{
		"cividis",
		"crest",
		"flare",
		"icefire",
		"inferno",
		"magma",
		"mako",
		"plasma",
		"rocket",
		"turbo",
		"viridis",
		"vlag",
	}
}

type Colormap struct {
	Palette palette.Palette
}

func (c Colormap) ToRGBA(x interface{}, a interface{}) color.RGBA {
	switch v := x.(type) {
	case nil:
		return color.RGBA{0, 0, 0, alpha(a)}
	case float64:
		r, g, b := rgbtouint8(c.interpolate_or_clip(v))
		return color.RGBA{r, g, b, alpha(a)}
	case uint8:
		r, g, b := rgbtouint8(c.Palette.Lookup(v))
		return color.RGBA{r, g, b, alpha(a)}
	default:
		fmt.Printf("warning: x (%T) must be float64 or uint8\n", x)
		return color.RGBA{0, 0, 0, alpha(a)}
	}
}

func (c Colormap) interpolate(x float64) [3]float32 {
	lut := c.Palette.LUT()

	x = math.Max(0.0, math.Min(1.0, x))
	a := int(x * 255.0)
	b := min(255, a+1)
	f := float32(x*255.0 - float64(a))

	return [3]float32{
		lut[a][0] + (lut[b][0]-lut[a][0])*f,
		lut[a][1] + (lut[b][1]-lut[a][1])*f,
		lut[a][2] + (lut[b][2]-lut[a][2])*f,
	}
}

func (c Colormap) interpolate_or_clip(x float64) [3]float32 {
	if x < 0.0 {
		return [3]float32{0.0, 0.0, 0.0}
	} else if x > 1.0 {
		return [3]float32{1.0, 1.0, 1.0}
	} else {
		return c.interpolate(x)
	}
}

func rgbtouint8(rgb [3]float32) (uint8, uint8, uint8) {
	return uint8(rgb[0] * 255.0), uint8(rgb[1] * 255.0), uint8(rgb[2] * 255.0)
}

func alpha(a interface{}) uint8 {
	switch v := a.(type) {
	case nil:
		return 255
	case float64:
		x := math.Max(0.0, math.Min(1.0, v))
		return uint8(x * 255.0)
	case int:
		return uint8(min(v, 255))
	case uint8:
		return v
	default:
		fmt.Printf("warning: alpha (%T) must be float64 or uint8\n", a)
		return 255
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
